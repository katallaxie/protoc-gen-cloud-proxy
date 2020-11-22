import * as cdk from '@aws-cdk/core'
import * as ec2 from '@aws-cdk/aws-ec2'
import * as ecs from '@aws-cdk/aws-ecs'
import * as iam from '@aws-cdk/aws-iam'
import * as elb from '@aws-cdk/aws-elasticloadbalancingv2'
import * as assets from '@aws-cdk/aws-ecr-assets'
import * as path from 'path'
import * as logs from '@aws-cdk/aws-logs'
import * as golang from 'aws-lambda-golang'

enum SubnetName {
  Public = 'Public',
  Application = 'Application'
}

export interface PlatformStackProps extends cdk.StackProps {
  name: string
  enhancedMonitoring?: boolean
}

export class Platform extends cdk.Stack {
  public readonly vpc: ec2.Vpc
  public readonly ecsCluster: ecs.Cluster
  public readonly image: cdk.DockerImageAssetLocation
  public readonly loadbalancer: elb.ApplicationLoadBalancer
  public readonly loadbalancerListener: elb.ApplicationListener
  public readonly loadbalancerTargetGroup: elb.NetworkTargetGroup
  public readonly taskDefinition: ecs.FargateTaskDefinition
  public readonly taskContainer: ecs.ContainerDefinition
  public readonly taskExecutionRole: iam.Role
  public readonly taskSecurityGroup: ec2.SecurityGroup
  public readonly taskRole: iam.Role
  public readonly service: ecs.FargateService
  public readonly dockerImage: assets.DockerImageAsset
  public readonly lambdaHandler: golang.GolangFunction

  constructor(scope: cdk.Construct, id: string, props: PlatformStackProps) {
    super(scope, id)

    // Creating a new VPC for Amazon MSK. It has three subnets.
    this.vpc = new ec2.Vpc(this, 'VPC', {
      cidr: '10.0.0.0/16',
      natGatewaySubnets: {
        subnetName: SubnetName.Public
      },
      subnetConfiguration: [
        {
          cidrMask: 24,
          name: SubnetName.Public,
          subnetType: ec2.SubnetType.PUBLIC
        },
        {
          cidrMask: 24,
          name: SubnetName.Application,
          subnetType: ec2.SubnetType.PRIVATE
        }
      ]
    })

    const image = new assets.DockerImageAsset(this, 'MyBuildImage', {
      directory: path.join(__dirname, '../../../')
    })

    // this.image = this.addDockerImageAsset({
    //   directoryName: path.join(__dirname, '../../../'),
    //   repositoryName: `${props.name}/proxy`,
    //   sourceHash: 'proxy-beta_2'
    // })

    // this.lambdaHandler = new golang.GolangFunction(this, 'test-function')

    this.taskExecutionRole = new iam.Role(this, 'TaskExecutionRole', {
      roleName: `${props.name}-proxy-execution`,
      managedPolicies: [
        iam.ManagedPolicy.fromManagedPolicyArn(
          this,
          'taskEcrReadOnly',
          'arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly'
        ),
        iam.ManagedPolicy.fromManagedPolicyArn(
          this,
          'taskCloudWatchLogs',
          'arn:aws:iam::aws:policy/CloudWatchLogsFullAccess'
        )
      ],
      assumedBy: new iam.ServicePrincipal('ecs-tasks.amazonaws.com', {})
    })

    this.taskRole = new iam.Role(this, 'TaskRole', {
      roleName: `${this.stackName}-proxy`,
      managedPolicies: [
        iam.ManagedPolicy.fromManagedPolicyArn(
          this,
          'taskAWSXRayDaemonWriteAccess',
          'arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess'
        ),
        iam.ManagedPolicy.fromManagedPolicyArn(
          this,
          'taskExecutionCloudWatchLogs',
          'arn:aws:iam::aws:policy/CloudWatchLogsFullAccess'
        )
      ],
      assumedBy: new iam.ServicePrincipal('ecs-tasks.amazonaws.com', {})
    })

    // ECS Cluster ...
    this.ecsCluster = new ecs.Cluster(this, 'Cluster', {})

    // ECS Task Security Group ...
    this.taskSecurityGroup = new ec2.SecurityGroup(this, 'SecurityGroup', {
      vpc: this.vpc,
      securityGroupName: `${props.name}-proxy`,
      description: `default proxy security`
    })
    this.taskSecurityGroup.addIngressRule(
      ec2.Peer.anyIpv4(),
      ec2.Port.tcp(8443)
    )
    this.taskSecurityGroup.addIngressRule(
      ec2.Peer.anyIpv4(),
      ec2.Port.tcp(9090)
    )

    // ECS Task ...
    this.taskDefinition = new ecs.FargateTaskDefinition(this, 'TaskDef', {
      memoryLimitMiB: 512,
      executionRole: this.taskExecutionRole,
      taskRole: this.taskRole
    })
    this.taskContainer = this.taskDefinition.addContainer('ProxyContainer', {
      essential: true,
      image: ecs.ContainerImage.fromRegistry(image.imageUri),
      environment: {
        PORT: '9090'
      },
      logging: new ecs.AwsLogDriver({
        streamPrefix: `${props.name}-proxy`,
        logGroup: new logs.LogGroup(this, 'LogGroup', {
          logGroupName: `${props.name}/ecs/proxy`,
          retention: logs.RetentionDays.ONE_WEEK
        })
      })
    })
    this.taskContainer.addPortMappings(
      ...[{ containerPort: 9090, protocol: ecs.Protocol.TCP }]
    )
    this.taskContainer.addPortMappings(
      ...[{ containerPort: 8443, protocol: ecs.Protocol.TCP }]
    )

    // ECS Service ...
    this.service = new ecs.FargateService(this, 'Service', {
      cluster: this.ecsCluster,
      taskDefinition: this.taskDefinition,
      securityGroup: this.taskSecurityGroup,
      vpcSubnets: { subnets: this.vpc.privateSubnets }
    })
    // this.loadbalancerTargetGroup.addTarget(this.service)

    // Setup AutoScaling policy ...
    const scaling = this.service.autoScaleTaskCount({ maxCapacity: 2 })
    scaling.scaleOnCpuUtilization('CpuScaling', {
      targetUtilizationPercent: 50,
      scaleInCooldown: cdk.Duration.seconds(60),
      scaleOutCooldown: cdk.Duration.seconds(60)
    })
  }
}
