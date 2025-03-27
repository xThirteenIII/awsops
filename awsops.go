package awsops

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// AWSOps provides AWS configurations along with AWS services clients.
type AWSOps struct {
    Config          aws.Config
    DynamoDBClient  *dynamodb.Client
}

// NewAWSOps returns a new empty AWSOps pointer.
func NewAWSOps() *AWSOps {
    return &AWSOps{}
}

// Copy returns a shallow copy of the AWSOps object.
func (a AWSOps) Copy() AWSOps {
    cp := a
    return cp
}
