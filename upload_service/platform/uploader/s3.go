package uploader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-yaml/yaml"
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

func uploadFileToS3(fileName string, fileContent []byte, description string) (string, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(conf.Region),

			CredentialsChainVerboseErrors: aws.Bool(true)},
		SharedConfigState: session.SharedConfigEnable,

		Profile: conf.Profile,
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return "", err
	}
	svc := s3.New(sess)

	bucket := conf.BucketName
	// This uploads the contents of the buffer to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:  aws.String(bucket),
		Key:     aws.String(fileName),
		Body:    bytes.NewReader(fileContent),
		Tagging: aws.String("Description=" + description),
	})
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return "", err
	}

	fmt.Println("File uploaded successfully!!!")
	return "Uploaded successfully", nil
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	var errors = false

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(100 << 20)

	fileName := r.FormValue("fileName")
	country := r.FormValue("country")
	description := r.FormValue("fileName")

	if fileName == "" {
		fmt.Println("Filename should not be empty")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Filename should not be empty")
		return
	}
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		fmt.Fprintf(w, "Error Retrieving the File")
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		errors = true
	}

	msg, err := uploadFileToS3(country+"/"+fileName, fileBytes, description)
	if err != nil {
		fmt.Fprint(w, err.Error())
		errors = true
	}
	// write this byte array to our temporary file
	// return that we have successfully uploaded our file!
	if !errors {
		fmt.Fprint(w, msg)
	}
}

func Handler(ctx *gin.Context) {
	uploadFile(ctx.Writer, ctx.Request)
}
