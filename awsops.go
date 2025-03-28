package awsops

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var generalConfig Config

func init() {
    var err error
    generalConfig, err = initConfig(context.TODO())
    if err != nil {
        log.Fatalf("\nerror loading aws configurations.")
    }
}

// AWSOps provides AWS services.
type AWSOps struct {
    DynamoDB  *DynamoService
}


// Config provides general configs.
type Config struct {
    awsConfig  aws.Config
}

func GetSDKConfig() aws.Config {
    return generalConfig.awsConfig
}

// NewAWSOps returns a new AWSOps pointer.
// The AWSOps structure is initialized.
//
// Returns the error encountered by initConfig.
func NewAWSOps() (*AWSOps, error) {

    return &AWSOps{}, nil
}

// Copy returns a shallow copy of the AWSOps object.
func (a AWSOps) Copy() AWSOps {
    cp := a
    return cp
}

// initConfig is the private function called by NewAWSOps.
//
// It returns the error encountered by the config.LoadDefaultConfig official AWS function.
// Optional configs are not supported yet.
func initConfig(ctx context.Context) (Config, error) {

    fmt.Printf("\nLoading AWS configurations... ")
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        return Config{}, fmt.Errorf("failed loading aws configurations. Here's why: %v", err)
    }
    fmt.Printf("Done\n")

    return Config{awsConfig: cfg}, nil
}
