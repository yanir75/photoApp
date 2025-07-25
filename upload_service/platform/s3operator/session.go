package s3operator

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-yaml/yaml"
)

type Config struct {
	Region     string `yaml:"Region"`
	BucketName string `yaml:"BucketName"`
	Profile    string `yaml:"Profile"`
}

type ObjectMetaData struct {
	Type         string
	Name         string
	Description  string
	ThumbnailUrl string
}

type Country struct {
	FirstImage string
	MetaData   []ObjectMetaData
}

type SafeManifest struct {
	mu sync.Mutex
	countryMap  map[string]*Country
}

var t time.Time = time.Now().Add(-time.Minute * 15)
var countryUrlTime map[string]time.Time = make(map[string]time.Time)

var objectFolderMap map[string]string = make(map[string]string)
var manifest SafeManifest = SafeManifest{countryMap:make(map[string]*Country)}
var countryUrlMap map[string][]ObjectMetaData = make(map[string][]ObjectMetaData)

var manifestDescription = "Manifest json of the countries and their first image key"
var s3ManifestKey string = "manifest.json"
var upload bool = true

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

var conf Config

func createManifest() {
	res := listObjects("", s3ManifestKey, 1)
	if len(res.Contents) == 0 {
		fmt.Println("Created manifest since it did not exist")
		uploadFileToS3(s3ManifestKey, []byte("{}"), manifestDescription,false)
	}
}

func loadManifest() {
	// fmt.Printf("test: %s\n",string(GetObject(s3ManifestKey)))
	err := json.Unmarshal(GetObject(s3ManifestKey), &manifest.countryMap)
	for key := range manifest.countryMap {
		countryUrlTime[key] = time.Now().Add(-time.Minute * 15)
		countryUrlMap[key] = []ObjectMetaData{}
	}
	if err != nil {
		fmt.Println("Coudln't load manifest: ", err)
	}
	// fmt.Println(manifest)
}

func updateManifest(country string, fileName string, fileType string, description string) {
	manifest.mu.Lock()
	if _, ok := manifest.countryMap[country]; !ok {
		manifest.countryMap[country] = &Country{"", []ObjectMetaData{}}
	}
	if fileType != "image" {
		manifest.countryMap[country].FirstImage = fileName + ".jpeg"
	} else {
		manifest.countryMap[country].FirstImage = fileName
	}
	manifest.countryMap[country].MetaData = append(manifest.countryMap[country].MetaData, ObjectMetaData{fileType, fileName, description, ""})
	countryUrlTime[country] = time.Now().Add(-time.Minute * 15)

	mapJson, err := json.Marshal(manifest.countryMap)
	if err != nil {
		fmt.Println("Couldn't marshal manifest: ", err)
	}
	uploadFileToS3(s3ManifestKey, mapJson, manifestDescription,false)
	manifest.mu.Unlock()

}

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
	createManifest()
	loadManifest()
}

func getSession() (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(conf.Region),

			CredentialsChainVerboseErrors: aws.Bool(true)},
		SharedConfigState: session.SharedConfigEnable,

		Profile: conf.Profile,
	})

	return sess, err
}

func generateS3Session() (*s3.S3, string) {
	sess, err := getSession()

	if err != nil {
		fmt.Println("Error creating session:", err)
		// return "", err
	}

	return s3.New(sess), conf.BucketName
}
