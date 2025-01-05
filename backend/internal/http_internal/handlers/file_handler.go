package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/OsqY/GoingNext/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type FileHandler struct {
	Config *config.Config
}

func NewFileHandler(config *config.Config) *FileHandler {
	return &FileHandler{Config: config}
}

func (f *FileHandler) SendFileToS3(w http.ResponseWriter, r *http.Request) {
	sess, err := session.NewSession(&aws.Config{
		Region:                        aws.String(f.Config.Aws.AwsRegion),
		Credentials:                   credentials.NewStaticCredentials(f.Config.Aws.AccessKeyId, f.Config.Aws.SecretAccessKey, ""),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		http.Error(w, "Error creating Aws Session", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	svc := s3.New(sess)

	key := fmt.Sprintf("user-images/%d-%s", time.Now().Unix(), header.Filename)

	var buf bytes.Buffer

	if _, err := io.Copy(&buf, file); err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(f.Config.Aws.S3Bucket),
		Key:    aws.String(key),

		Body: bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		log.Printf("Error uploading to S3: %v", err)
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", f.Config.Aws.S3Bucket, f.Config.Aws.AwsRegion, key)
	json.NewEncoder(w).Encode(response)
}
