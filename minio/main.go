package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	endPoint := "localhost:9000"
	accessKey := "jake"
	secretKey := "jake1234"

	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
		Trace:              nil,
		Region:             "",
		BucketLookup:       0,
		CustomRegionViaURL: nil,
		TrailingHeaders:    false,
		CustomMD5:          nil,
		CustomSHA256:       nil,
	})
	if err != nil {
		log.Fatalln("Error initializing MinIO client:", err)
	}

	log.Println("download start")
	object, err := minioClient.GetObject(context.Background(), "test", "ai+6.mp4", minio.GetObjectOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	defer object.Close()
	log.Println("download end")

	localFile, err := os.Create("/home/jake/minio/z.mp4")
	if err != nil {
		log.Println(err.Error())
	}
	defer localFile.Close()

	log.Println("copy start")

	if _, err = io.Copy(localFile, object); err != nil {
		log.Println(err.Error())
	}
	log.Println("copy end")
}
