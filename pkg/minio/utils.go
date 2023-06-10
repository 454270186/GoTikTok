package minio

import (
	"context"
	"errors"
	"log"

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