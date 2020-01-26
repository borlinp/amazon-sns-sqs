package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/borlinp/amazon-sns-sqs/common"
)

func main() {
	awsSession := common.BuildSession()
	if awsSession == nil {
		os.Exit(1)
	}

	svc := sqs.New(awsSession, nil)
	if svc == nil {
		log.Println("There was a problem getting a new SQS instance")
		os.Exit(1)
	}

	listQueuesOutput, err := svc.ListQueues(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	if listQueuesOutput == nil || len(listQueuesOutput.QueueUrls) == 0 {
		log.Println("There are no queues to display.")
		return
	}

	for _, queue := range listQueuesOutput.QueueUrls {
		if queue == nil {
			log.Println("<empty>")
			continue
		}
		fmt.Println(" ", *queue, "  ")
	}
}
