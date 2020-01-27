package common

type Creds struct {
	AccessKey string
	SecretKey string
	Region    string
}

func GetCredentials() Creds {
	return Creds{
		AccessKey: "", //Load your access key here
		SecretKey: "", //Load your secret key here
		Region:    "", //Load your region here
	}
}
