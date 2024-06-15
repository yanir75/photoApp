package s3operator

import (
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func urlSigner(objectName string, timeInMinutesToSign ...int) string {
	timeInMinutes := 15
	if len(timeInMinutesToSign) > 0 {
		timeInMinutes = timeInMinutesToSign[0]
	}

	svc, bucket := generateS3Session()

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectName),
	})
	urlStr, err := req.Presign(time.Duration(timeInMinutes) * time.Minute)

	if err != nil {
		fmt.Println("Failed to sign request", err)
	}
	// fmt.Println(urlStr)
	return urlStr
}

func listObjects(delimeter string, prefix string, maxKeys int64) *s3.ListObjectsV2Output {
	svc, bucket := generateS3Session()
	li, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Delimiter: aws.String(delimeter),
		Prefix:    aws.String(prefix),
		MaxKeys:   aws.Int64(maxKeys),
	})
	if err != nil {
		fmt.Println("Error listing folders:", err)
	}

	return li
}

func GetObject(objectName string) []byte {
	svc, bucket := generateS3Session()
	res, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3ManifestKey),
	})
	if err != nil {
		fmt.Println("Error getting object:", err)
	}

	defer res.Body.Close()
	fileBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Couldn't read s3 file:", err)
	}
	if err != nil {
		fmt.Println("Couldn't close s3 bodyfile", err)
	}
	return fileBytes

}

// func getS3Folders() ([]*s3.CommonPrefix){
// 	li := listObjects("/","",1000)

// 	return li.CommonPrefixes
// }

// func mapItemsToFolder()(map[string]string) {

// 	objectFolderMap = make(map[string]string)

// 	for _,folder:= range getS3Folders() {
// 		li := listObjects("",*folder.Prefix,1)
// 		objectFolderMap[(*folder.Prefix)[:len(*folder.Prefix)-1]] = *li.Contents[0].Key
// 	}

// 	return objectFolderMap
// }

func GenerateUrlMap() map[string]string {
	if time.Since(t).Minutes() > 10 || upload {
		t = time.Now()
		// objectFolderMap = mapItemsToFolder()
		for key, value := range manifest.countryMap {

			url := urlSigner(value.FirstImage, 15)
			objectFolderMap[key] = url
		}
		// fmt.Println(time.Since(t).Minutes())
		// fmt.Println(objectFolderMap)
	}
	upload = false
	return objectFolderMap
}

func GenerateUrlCountryMap(country string) map[string][]ObjectMetaData {
	if time.Since(countryUrlTime[country]).Minutes() > 10 {
		countryUrlMap[country] = []ObjectMetaData{}
		countryUrlTime[country] = time.Now()
		// objectFolderMap = mapItemsToFolder()
		for _, value := range manifest.countryMap[country].MetaData {
			url := urlSigner(value.Name, 15)
			thumbnailUrl := url
			if value.Type != "image" {
				thumbnailUrl = urlSigner(value.Name + ".jpeg")
			}
			countryUrlMap[country] = append(countryUrlMap[country], ObjectMetaData{value.Type, url, value.Description, thumbnailUrl})
		}
		// fmt.Println(time.Since(t).Minutes())
		// fmt.Println(objectFolderMap)
	}
	return map[string][]ObjectMetaData{country: countryUrlMap[country]}
}
