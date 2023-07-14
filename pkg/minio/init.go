package minio

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient          *minio.Client
	MinioEndpoint        string
	MinioAccessKeyId     string
	MinioSecretAccessKey string
	MinioUseSSL          = false
	VideoBucketName      string
)

// initialize minio storage object
func init() {
	// init minio env
	mEnv, err := godotenv.Read()
	if err != nil {
		panic(err)
	}
	MinioEndpoint = mEnv["M_ENDPOINT"]
	MinioAccessKeyId = mEnv["M_ACCESS_KEY_ID"]
	MinioSecretAccessKey = mEnv["M_SECRET_ACCESS_KEY"]
	VideoBucketName = mEnv["M_VIDEO_BUCKET_NAME"]

	// connect to minio server
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
