package common

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func BuildSession() *session.Session {

	creds := GetCredentials()

	sessionConfig := &aws.Config{
		Region:      aws.String(creds.Region),
		Credentials: credentials.NewStaticCredentials(creds.AccessKey, creds.SecretKey, ""),
	}

	sess, err := session.NewSession(sessionConfig)
	if err != nil {
		log.Println("error establishing session")
		return nil
	}
	return sess

}
