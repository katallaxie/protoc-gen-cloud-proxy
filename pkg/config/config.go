package config

type AWSRegion string

const (
	AWSRegion_EU_West_1 AWSRegion = "eu-west-1"
)

type Config struct {
	AWSRegion AWSRegion
}
