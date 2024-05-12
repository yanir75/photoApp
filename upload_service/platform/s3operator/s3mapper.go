package s3operator
import (
	"fmt"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"

)

func GetS3Folders(){
	sess,err := getSession()

	if err != nil {
		fmt.Println("Error creating session:", err)
		// return "", err
	}
	
	svc := s3.New(sess)

	bucket := conf.BucketName

	li,err :=svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Delimiter: aws.String("/"),
		Prefix: aws.String(""),
	})
	if err != nil {
		fmt.Println("Error listing folders:", err)
	}
	fmt.Println(li.CommonPrefixes[0].String())
}

func UrlSigner(objectName string , timeInMinutesToSign ... int) (string){
	sess,err := getSession()
	timeInMinutes := 15
	if len(timeInMinutesToSign) > 0 {
		timeInMinutes = timeInMinutesToSign[0]
	}
	if err != nil {
		fmt.Println("Error creating session:", err)
		// return "", err
	}
	svc := s3.New(sess)

	bucket := conf.BucketName
    req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(objectName),
    })
    urlStr, err := req.Presign(time.Duration(timeInMinutes) * time.Minute)

    if err != nil {
        fmt.Println("Failed to sign request", err)
    }
    return urlStr
}