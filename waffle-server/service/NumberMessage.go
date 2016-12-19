package service

type InsertRequest struct {
	User   string `json:"user"`
	Number int    `json:"number"`
}

type InsertResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
