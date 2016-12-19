package endpoint

import (
	"github.com/1ambda/gokit-waffle/waffle-server/message"
	"github.com/1ambda/gokit-waffle/waffle-server/service"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func CreateGenerationEndpoint(svc service.NumberService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(message.GenerateRequest)

		msg, err := svc.Generate(req.User, req.Number)
		res := message.GenerateResponse{
			Message: msg, Error: "",
		}

		if err != nil {
			res.Error = err.Error()
			return res, nil
		}

		return res, nil
	}
}
