package service

type HasError interface {
	error() error
}

type ErrResponse struct {
	Error error `json:"error,omitempty"`
}

func (r ErrResponse) error() error {
	return r.Error
}

func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{Error: err}
}
