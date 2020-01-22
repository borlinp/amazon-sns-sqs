package common

import "os"

type Creds struct {
	AccessKey string
	SecretKey string
	Region    string
}

func GetCredentials() Creds {
	return Creds{
		AccessKey: os.Getenv("ACCESS_KEY"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Region:    os.Getenv("REGION"),
	}
}
