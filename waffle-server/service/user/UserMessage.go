package user

import "github.com/1ambda/gokit-waffle/waffle-server/service/common"

type UserListRequest struct {
}

type UserListResponse struct {
	Users []common.User
	common.ErrResponse
}

type UserRequest struct {
	User string
}

type UserResponse struct {
	User  string `json:"user"`
	Total int    `json:"total"`
	common.ErrResponse
}
