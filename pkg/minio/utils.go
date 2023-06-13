package minio

import (
	"context"
	"errors"
	"io"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

func CreateBucket(bucketName string) error {
	if len(bucketName) <= 0 {
		return errors.New("invalid bucket name")
	}

	location := "beijing"
	ctx := context.Background()

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		isExist, bucketErr := minioClient.BucketExists(ctx, bucketName)
		if bucketErr == nil && isExist {
			log.Printf("Bucket %s has existed\n", bucketName)
			return nil
		} else {
			return err
		}
	}

	return nil
}

// Upload local file
func UploadLocalFile(buckName string, objectName string, filePath string, contentType string) (int64, error) {
	ctx := context.Background()
	object, err := minioClient.FPutObject(ctx, buckName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println("upload failed: ", err.Error())
		return 0, err
	}

	log.Printf("Successfully upload %s of size %d\n", objectName, object.Size)
	return object.Size, nil
}

// Upload file
func UploadFile(bucketName string, objectName string, reader io.Reader, objectSize int64) error {
	ctx := context.Background()
	n, err := minioClient.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		log.Println("upload failed: ", err.Error())
		return err
	}
	
	log.Printf("Successfully upload %s of bytes %d\n", objectName, n.Size)
	return nil
}

func GetFileURL(bucketName string, fileName string, exp time.Duration) (*url.URL, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	if exp < 1 {
		exp = time.Second * 60 * 60 * 24
	}

	preSignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, exp, reqParams)
	if err != nil {
		log.Printf("get url of file %s from bucket %s failed: %s\n", fileName, bucketName, err.Error())
		return nil, err
	}

	return preSignedUrl, nil
}