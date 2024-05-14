package s3operator

import (
	"fmt"

	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)
var t time.Time = time.Now().Add(- time.Minute * 15)
var objectFolderMap map[string]string = make(map[string]string)

func urlSigner(objectName string , timeInMinutesToSign ... int) (string){
	timeInMinutes := 15
	if len(timeInMinutesToSign) > 0 {
		timeInMinutes = timeInMinutesToSign[0]
	}

	svc,bucket := generateS3Session()

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

func listObjects(delimeter string, prefix string,maxKeys int64)(*s3.ListObjectsV2Output){
	svc,bucket := generateS3Session()
	li,err :=svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Delimiter: aws.String(delimeter),
		Prefix: aws.String(prefix),
		MaxKeys:  aws.Int64(maxKeys),
	})
	if err != nil {
		fmt.Println("Error listing folders:", err)
	}
	
	return li
}

func getS3Folders() ([]*s3.CommonPrefix){
	li := listObjects("/","",1000)

	return li.CommonPrefixes
}

func mapItemsToFolder()(map[string]string) {


	objectFolderMap = make(map[string]string)

	for _,folder:= range getS3Folders() {
		li := listObjects("",*folder.Prefix,1)
		objectFolderMap[(*folder.Prefix)[:len(*folder.Prefix)-1]] = *li.Contents[0].Key
	}
	
	return objectFolderMap
}

func GenerateUrlMap()(map[string]string){
	if (time.Since(t).Minutes() > 10){
		t = time.Now()
		objectFolderMap = mapItemsToFolder()
		for key,value := range objectFolderMap {
			objectFolderMap[key] = urlSigner(value,15)
		}
	// fmt.Println(time.Since(t).Minutes())
		
	}
	return objectFolderMap
}


