package user

import . "github.com/1ambda/gokit-waffle/waffle-server/service/common"

type UserListRequest struct {
}

type UserListResponse struct {
	Users []User
	ErrResponse
}
