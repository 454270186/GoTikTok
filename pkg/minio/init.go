package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client
	// MinioEndpoint        = "172.20.10.2:9000"
	// MinioEndpoint        = "192.168.2.44:9000"
	MinioEndpoint        = "10.14.13.212:9000"
	MinioAccessKeyId     = "tiktokMinio"
	MinioSecretAccessKey = "tiktokMinio"
	MinioUseSSL          = false
	VideoBucketName      = "tiktok-video"
)

// initialize minio storage object
func init() {
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Secure: MinioUseSSL,
		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln("minio client init failed")
	}

	minioClient = client
	if err := CreateBucket(VideoBucketName); err != nil {
		log.Fatalln("bucket create failed")
	}
	log.Println("------------------------------")
	log.Println("Minio client init successfully")
	log.Println("------------------------------")
}
