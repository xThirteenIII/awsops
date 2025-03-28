package awsops

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// InitDynamoDBClient creates a new DynamoDBClient
func (a *AWSOps) InitDynamoDBClient() {
   a.DynamoDBClient = dynamodb.NewFromConfig(a.cfg.AWSConfig)
}

// GetMarshaledKey returns a key ready to be used for dynamoDB operations. 
// Every database operation that needs an ItemInput{} requires a key in the map[string]types.AttributeValue{} format.
// Exits with Fatalf if marshaling is unsuccessful.
//
// E.g. 
// dynamodb.DeleteItemInput{Key: GetMarshaledKey("myKeyValue", "myKeyName")} is needed by the dynamodb.Client.DeleteItem() operation.
func GetMarshaledKey(keyValue, keyName string) map[string]types.AttributeValue {
    marshaledKey, err := attributevalue.Marshal(keyValue)
    if err != nil {
        log.Fatalf("\nerror marshaling %s key. Here's why: %v", err)
    }
    return map[string]types.AttributeValue{keyName: marshaledKey}
}

