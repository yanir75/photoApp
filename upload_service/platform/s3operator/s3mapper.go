package s3operator
import (
	"fmt"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Folders(){
	sess,err := getSession()

	if err != nil {
		fmt.Println("Error creating session:", err)
		// return "", err
	}
	
	svc := s3.New(sess)

	bucket := conf.BucketName

	li,_ :=svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Delimiter: aws.String("/"),
		Prefix: aws.String(""),
	})
	fmt.Println(li.CommonPrefixes)
}