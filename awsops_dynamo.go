package awsops

import (

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// InitDynamoDBClient creates a new DynamoDBClient
func (a *AWSOps) InitDynamoDBClient() {
   a.DynamoDBClient = dynamodb.NewFromConfig(a.cfg.AWSConfig)
}
