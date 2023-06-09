package logic

import (
	"context"
	"errors"
	"strconv"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/rpc/publish/internal/svc"
	"github.com/454270186/GoTikTok/rpc/publish/types/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

var PublishDB dal.PublishDB
var UserDB dal.UserDB

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *publish.PublishListReq) (*publish.PublishListRes, error) {
	// todo: add your logic here and delete this line
	userID, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return &publish.PublishListRes{StatusCode: -1}, err
	}

	videoList, err := PublishDB.GetListByID(l.ctx, uint(userID))
	if err != nil {
		return nil, err
	}

	respVideoList := make([]*publish.Video, 0)
	for _, video := range videoList {
		author, err := UserDB.GetById(l.ctx, video.AuthorID)
		if err != nil {
			return &publish.PublishListRes{StatusCode: -1}, err
		}

		respVideo := publish.Video{
			Id:            int64(video.ID),
			Author:        convertUser(author),
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
		}

		respVideoList = append(respVideoList, &respVideo)
	}

	if respVideoList == nil {
		return &publish.PublishListRes{StatusCode: -1}, errors.New("empty video list")
	}

	return &publish.PublishListRes{
		StatusCode: 0,
		VideoList:  respVideoList,
	}, nil
}

// Convert dal.User model to RPC User model
func convertUser(dalUser *dal.User) *publish.User {
	if dalUser == nil {
		return nil
	}

	return &publish.User{
		Id:            int64(dalUser.ID),
		Name:          dalUser.Username,
		FollowCount:   dalUser.FollowingCount,
		FollowerCount: dalUser.FollowerCount,
		IsFollow:      true,
	}
}
