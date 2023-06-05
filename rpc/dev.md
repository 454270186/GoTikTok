## user rpc service

### Register
**Request body**
```go
type RegisterReq struct {
	Username string
	Password string
}
```

**Response body**
```go
type RegisterRes struct {
	StatusCode int32
	UserId     string
    Token        string
	RefreshToken string
}
```

**rpc interface**
```go
Register(in *user.RegisterReq) (*user.RegisterRes, error)
```

### Login
**Request body**
```go
type LoginReq struct {
	Username string
	Password string
}
```

**Response body**
```go
type LoginRes struct {
	StatusCode int32
	UserId     string
}
```

**rpc interface**
```go
Login(in *user.LoginReq) (*user.LoginRes, error)
```

### Get User by ID
**User struct**
```go
type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}
```

**Request body**
```go
type GetUserReq struct {
	UserId string
}
```

**Response body**
```go
type GetUserRes struct {
	StatusCode int32
	User       *User
}
```

**rpc interface**
```go
GetUserById(in *user.GetUserReq) (*user.GetUserRes, error)
```

### Get new token by refresh
**Request body**
```go
type RefreshReq struct {
	Token        string
	RefreshToken string
}
```

**Response body**
```go
type RefreshRes struct {
	StatusCode int32
	Token      string
}
```

**rpc interface**
```go
Refresh(in *user.RefreshReq) (*user.RefreshRes, error)
```

<br>
<br>



## Publish rpv service

### User and Video struct
```go
type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type Video struct {
	Id            int64
	Author        *User
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
}
```
### Publish list

**Request body**
```go
type PublishListReq struct {
	UserId string
}
```

**Response body**
```go
type PublishListRes struct {
	StatusCode int32
	VideoList  []*Video
}
```

**rpc interface**
```go
PublishList(in *publish.PublishListReq) (*publish.PublishListRes, error)
```

### Publish action

**Request body**
```go
type PublishActionReq struct {
	Data  []byte
	Title string
}
```

**Response body**
```go
type PublishActionRes struct {
	StatusCode int32
}
```

**rpc interface**
```go
PublishAction(in *publish.PublishActionReq) (*publish.PublishActionRes, error)
```

<br>
<br>



## Feed rpc service

### Get user feed
**Request body**
```go
type FeedReq struct {
	LastestTime int64
}
```

**Response body**
```go
type FeedRes struct {
	StatusCode int32
	VideoList  []*Video
	NextTime   int64
}
```

**rpc interface**
```go
GetUserFeed(in *feed.FeedReq) (*feed.FeedRes, error)
```
