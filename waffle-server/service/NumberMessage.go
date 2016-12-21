package service

type InsertRequest struct {
	User   string `json:"user"`
	Number int    `json:"number"`
}

type InsertResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}

func (r InsertResponse) error() error { return r.Error }

type QueryRequest struct {
	User string
}

type QueryResponse struct {
	User  string `json:"user"`
	Total int    `json:"total"`
	Error error  `json:"error,omitempty"`
}

func (r QueryResponse) error() error { return r.Error }
