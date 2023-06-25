package pack

import (
	"context"
	"errors"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/pkg/auth"
	"github.com/454270186/GoTikTok/rpc/user/types/user"
)

var userDB dal.UserDB
var ctx = context.Background()

func init() {
	userDB = dal.NewUserDB()
}

func IsUserExist(username string) (uint, bool, error) {
	userID, isExist, err := userDB.IsUserExist(ctx, username)
	if err != nil {
		return 0, false, err
	}

	return userID, isExist, nil
}

func GetUserByID(userID uint64) (*user.User, error) {
	dalUser, err := userDB.GetById(ctx, uint(userID))
	if err != nil {
		return nil, err
	}

	return &user.User{
		Id: int64(dalUser.ID),
		Name: dalUser.Username,
		FollowCount: dalUser.FollowingCount,
		FollowerCount: dalUser.FollowerCount,
		IsFollow: false,
	}, nil
}

func CreateNewUser(username string, password string) (uint, error) {
	_, isExist, err := userDB.IsUserExist(ctx, username)
	if err != nil {
		return 0, err
	} else if isExist {
		return 0, errors.New("user has existed")
	}
	
	newUser := &dal.User{
		Username: username,
		Password: password,
	}

	err = userDB.CreateUser(ctx, newUser)
	if err != nil {
		return 0, err
	}

	return newUser.ID, nil
}

func CheckPassword(userID uint, inPassword string) bool {
	dalUser, err := userDB.GetById(ctx, userID)
	if err != nil {
		return false
	}

	if isPwdOK := auth.ComparePwd(dalUser.Password, inPassword); !isPwdOK {
		return false
	}

	return true
}