package message

type GenerateRequest struct {
	User   string `json:"user"`
	Number int    `json:"number"`
}

type GenerateResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
