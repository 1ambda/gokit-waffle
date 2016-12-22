package number

import "github.com/1ambda/gokit-waffle/waffle-server/service/common"

type InsertRequest struct {
	User string `json:"user"`
	// TODO: type alias
	Number int `json:"number"`
}

type InsertResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
	common.ErrResponse
}

type QueryRequest struct {
	User string
}

type QueryResponse struct {
	User  string `json:"user"`
	Total int    `json:"total"`
	common.ErrResponse
}