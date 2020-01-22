package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/sns"

	"jabbok.ci.org/borlinp/amazon-sqs-sns/common"
)

func main() {
	awsSession := common.BuildSession()
	if awsSession == nil {
		os.Exit(1)
	}

	svc := sns.New(awsSession, nil)
	if svc == nil {
		log.Println("There was a problem getting a new SNS instance")
		os.Exit(1)
	}

	listTopicsOutput, err := svc.ListTopics(nil)
	if err != nil {
		fmt.Println("Unable to get topics: ", err)
		os.Exit(1)
	}
	if listTopicsOutput == nil || len(listTopicsOutput.Topics) == 0 {
		log.Println("There are no queues to display.")
		return
	}

	for _, queue := range listTopicsOutput.Topics {
		if queue == nil {
			log.Println("<empty>")
			continue
		}
		fmt.Println(*queue)
	}
}
