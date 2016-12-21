package service

type InsertRequest struct {
	User   string `json:"user"`
	Number int    `json:"number"`
}

type InsertResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}

type QueryRequest struct {
	User string
}

type QueryResponse struct {
	User  string `json:"user"`
	Total int    `json:"total"`
	Error error  `json:"error,omitempty"`
}
