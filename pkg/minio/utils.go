package minio

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	ffmpeg "github.com/u2takey/ffmpeg-go"
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
func UploadFile(bucketName string, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	ctx := context.Background()
	n, err := minioClient.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
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
	if exp < 1 {
		exp = time.Second * 60 * 60 * 24
	}

	preSignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, exp, nil)
	if err != nil {
		log.Printf("get url of file %s from bucket %s failed: %s\n", fileName, bucketName, err.Error())
		return nil, err
	}

	return preSignedUrl, nil
}

// UploadVideo() uploads video and cover data to minio bucket
// Encapsulates UploadFile() for video and cover
func UploadVideo(bucketName, videoFileName, coverFileName string, videoData []byte) error {
	videoReader := bytes.NewReader(videoData)
	
	// upload video data
	err := UploadFile(bucketName, videoFileName, videoReader, int64(len(videoData)), "video/mp4")
	if err != nil {
		return errors.New("video upload failed: " + err.Error())
	}

	playUrl, err := GetFileURL(bucketName, videoFileName, 0)
	if err != nil {
		return err
	}

	coverData, err := getOneFrameAsJpeg(playUrl.String())
	if err != nil {
		return errors.New("ffmpeg error: " + err.Error())
	}

	// upload cover
	coverReader := bytes.NewReader(coverData)
	err = UploadFile(bucketName, coverFileName, coverReader, int64(len(coverData)), "image/jpeg")
	if err != nil {
		return errors.New("cover upload failed: " + err.Error())
	}

	return nil
}

// 从视频流中截取一帧作为封面
func getOneFrameAsJpeg(playUrl string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	
	err := ffmpeg.Input(playUrl).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		   		  Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
				  WithOutput(reader, os.Stdout).Run()
	if err != nil {
		return nil, errors.New("ffmpeg failed: " + err.Error())
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, errors.New("image decode failed")
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), nil
}