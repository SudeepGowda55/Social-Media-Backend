package models

import "image"

type CreateUser struct {
	Email        string `json:"email"`
	UserName     string `json:"userName"`
	ProfileName  string `json:"profileName"`
	Age          string `json:"age"`
	Place        string `json:"place"`
	CoverPhoto   string `json:"coverPhoto"`
	ProfilePhoto string `json:"profilePhoto"`
}

type UserDetailsQuery struct {
	UserId string `json:"userId"`
	DID    string `json:"did"`
}

type UserDetails struct {
	Email        string `json:"email"`
	UserName     string `json:"userName"`
	ProfileName  string `json:"profileName"`
	CoverPhoto   string `json:"coverPhoto"`
	ProfilePhoto string `json:"profilePhoto"`
}

type FollowDetails struct {
	UserId      string   `json:"userId"`
	FollowersNo int16    `json:"followersNo"`
	FollowingNo int16    `json:"followingNo"`
	Followers   []string `json:"followers"`
	Following   []string `json:"following"`
}

type Post struct {
	PostId      string      `json:"postId"`
	UserId      string      `json:"userId"`
	PostContent string      `json:"postContent"`
	CreatedAt   string      `json:"createdAt"`
	UpdatedAt   string      `json:"updatedAt"`
	Image       image.Image `json:"image"`
}

type Comments struct {
	CommentId   string `json:"commentId"`
	PostId      string `json:"postId"`
	UserId      string `json:"userId"`
	Comment     string `json:"comment"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	CommentedBy string `json:"commentedBy"`
}
