package minio

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	CreateBucket("tiktoktest")
}

func TestUploadLocalFile(t *testing.T) {
	info, err := UploadLocalFile("tiktoktest", "test.mp4", "./test.mp4", "video/mp4")
	log.Println(info, err)
}

func TestUploadFile(t *testing.T) {
	file, _ := os.Open("./test.mp4")
	defer file.Close()

	fileInfo, _ := os.Stat("./test.mp4")

	err := UploadFile("tiktoktest", "test.mp4", file, fileInfo.Size())
	log.Println(err)
}

func TestGetFileURL(t *testing.T) {
	url, err := GetFileURL("tiktoktest", "test.mp4", 0)
	log.Println(url, err, strings.Split(url.String(), "?")[0])
	log.Println(url.Path, url.RawPath)
}