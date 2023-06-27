package logic

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/pkg/minio"
	"github.com/454270186/GoTikTok/rpc/publish/internal/svc"
	"github.com/454270186/GoTikTok/rpc/publish/types/publish"

	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *publish.PublishActionReq) (*publish.PublishActionRes, error) {
	MinioVideoBucketName := minio.VideoBucketName
	videoData := in.Data

	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// upload video
	fileName := u2.String() + "." + "mp4"
	err = minio.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)), "video/mp4")
	if err != nil {
		// return nil, err
		return nil, errors.New("minio.UploadFile")
	}

	// get video url
	url, err := minio.GetFileURL(MinioVideoBucketName, fileName, 0)
	if err != nil {
		// return nil, err
		return nil, errors.New("minio.GetFileURL")
	}

	// get cover from video stream
	u3, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	coverData, err := getOneFrameAsJpeg(url.String())
	if err != nil {
		return nil, errors.New("ffmpeg error: " + err.Error())
	}

	// upload cover
	coverPath := u3.String() + "." + "jpeg"
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadFile(MinioVideoBucketName, coverPath, coverReader, int64(len(coverData)), "image/jpeg")
	if err != nil {
		return nil, err
	}

	// get cover url
	coverURL, err := minio.GetFileURL(MinioVideoBucketName, coverPath, 0)
	if err != nil {
		return nil, err
	}

	err = pack.CreateVideo(uint(in.Uid), url.String(), coverURL.String(), in.Title)
	if err != nil {
		return nil, err
	}

	return &publish.PublishActionRes{
		StatusCode: 0,
	}, nil
}

// 从视频流中截取一帧作为封面
func getOneFrameAsJpeg(playUrl string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	log.Println(playUrl)
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