package pack

import (
	"context"

	"github.com/454270186/GoTikTok/dal"
)

var ctx = context.Background()

// Init database connection for all tables
func init() {
	userDB = dal.NewUserDB()
	publishDB = dal.NewPublishDB()
	feedDB = dal.NewFeedDB()
}