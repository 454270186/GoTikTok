package logic

import (
	// "fmt"
	"fmt"
	"log"
	"os"

	// "os"
	"testing"
)

func TestGetOneFrameAsJpeg(t *testing.T) {
	path := os.Getenv("PATH")
	ffmpegPath := "/Users/yuerfei/ffmpeg_exe/bin"
	newPath := fmt.Sprintf("%s:%s", path, ffmpegPath)
	err := os.Setenv("PATH", newPath)
	if err != nil {
		log.Println("无法设置环境变量")
		return
	}
	reader, err := getOneFrameAsJpeg("./sst.mp4")
	if err != nil {
		log.Println(err, "errerrerrerr")
	}
	log.Println(reader, "ssss")
}
