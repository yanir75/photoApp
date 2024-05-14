package s3operator

import (
	"fmt"

	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-yaml/yaml"
	"github.com/aws/aws-sdk-go/service/s3"

	
)


type Config struct {
	Region     string `yaml:"Region"`
	BucketName string `yaml:"BucketName"`
	Profile    string `yaml:"Profile"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

var conf Config

func init() {
	f, err := os.ReadFile("go_config.yaml")
	if err != nil {
		processError(err)
	}
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		processError(err)
	}
	fmt.Println("Loaded successfully")

}

func getSession() (*session.Session,error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(conf.Region),

			CredentialsChainVerboseErrors: aws.Bool(true)},
		SharedConfigState: session.SharedConfigEnable,

		Profile: conf.Profile,
	})

	return sess,err
}

func generateS3Session()(*s3.S3,string){
	sess,err := getSession()

	if err != nil {
		fmt.Println("Error creating session:", err)
		// return "", err
	}

	return s3.New(sess),conf.BucketName
}