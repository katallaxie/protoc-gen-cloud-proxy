#!/usr/bin/env node
import 'source-map-support/register'
import * as cdk from '@aws-cdk/core'
import { Platform } from '../lib'

const app = new cdk.App()

// extract context variables for the stacks, or default to generated one
const name = app.node.tryGetContext('projectName')

// creating the platform
const platform = new Platform(app, 'Platform', { name })

app.synth()
