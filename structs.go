package awsops

import (
	"fmt"

)

type awsService interface{
    initClient()
    getServiceName() string
}

func initServiceClient(s awsService) {

    fmt.Printf("\nInitializing %s Client...", s.getServiceName())
    s.initClient()
    fmt.Printf("Done\n")
}
