package s3operator

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)



func uploadFileToS3(fileName string, fileContent []byte, description string) (string, error) {
	sess,err := getSession()

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
		return "Not Uploaded successfully", err
	}

	fmt.Println("File uploaded successfully!!!")
	return "Uploaded successfully", nil
}

func uploadFile(ctx *gin.Context) {
	fmt.Println("File Upload Endpoint Hit")
	var errors = false
	r := ctx.Request
	w := ctx.Writer
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
	fhs := r.MultipartForm.File["myFiles"]
	msg := ""
	for index, handler := range fhs {
		file, err := handler.Open()
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
		types := strings.Split(handler.Header.Get("Content-Type"), "/")
		fileType, ending := types[0],types[1]
		key := fmt.Sprintf("%s/%s-%d.%s",country,fileName,index,ending)
		if fileType == "image"{
			go updateManifest(country,key)
		}
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

		message, err := uploadFileToS3(key, fileBytes, description)
		if err != nil {
			fmt.Fprint(w, err.Error())
			msg += fmt.Sprintf("File number %d: was not %s \\n", index, message)
		} else {
			msg += fmt.Sprintf("File number %d: was %s \\n", index, message)
			upload = true
		}
	}
	// write this byte array to our temporary file
	// return that we have successfully uploaded our file!
	if !errors {
		ctx.Redirect(http.StatusSeeOther, "/upload?msg="+msg)
	}
}

func Handler(ctx *gin.Context) {
	uploadFile(ctx)
}