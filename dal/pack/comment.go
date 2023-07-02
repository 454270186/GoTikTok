package pack

import (
	"strconv"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/rpc/comment/comments"
)

var commentDB dal.CommentDB

func GetCommentByID(commentID uint) (*comments.Comment, error) {
	dalComment, err := commentDB.GetById(ctx, commentID)
	if err != nil {
		return nil, err
	}

	return &comments.Comment{
		Id: int64(dalComment.ID),
		User: getRpcUser(dalComment.UserID),
		Content: dalComment.Content,
		CreateDate: dalComment.CreatedAt.Format("01-02"),
	}, nil
}

func GetCommentByVideoID(videoIDstr string) ([]*comments.Comment, error) {
	videoID, err := strconv.ParseUint(videoIDstr, 10, 64)
	if err != nil {
		return nil, err
	}

	dalComments, err := commentDB.GetByVideoID(ctx, uint(videoID))
	if err != nil {
		return nil, err
	}
	rpcComments := []*comments.Comment{}
	for _, dalComment := range dalComments {
		rpcComment := comments.Comment{
			Id: int64(dalComment.ID),
			User: getRpcUser(dalComment.UserID),
			Content: dalComment.Content,
			CreateDate: dalComment.CreatedAt.Format("01-02"),
		}

		rpcComments = append(rpcComments, &rpcComment)
	}

	return rpcComments, nil
}

func AddComment(videoIDstr, userIDstr, text string) (uint, error) {
	videoID, err := strconv.ParseUint(videoIDstr, 10, 64)
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		return 0, err
	}

	return commentDB.Add(ctx, uint(videoID), uint(userID), text)
}

func DelComment(commentIDstr string) error {
	commentID, err := strconv.ParseUint(commentIDstr, 10, 64)
	if err != nil {
		return err
	}

	return commentDB.Del(ctx, uint(commentID))
}

func getRpcUser(userID uint) *comments.User {
	dalUser, err := userDB.GetById(ctx, userID)
	if err != nil {
		return &comments.User{}
	}

	return &comments.User{
		Id: int64(dalUser.ID),
		Name: dalUser.Username,
		FollowCount: dalUser.FollowingCount,
		FollowerCount: dalUser.FollowerCount,
		IsFollow: true,
	}
}