package service

import (
	"github.com/1ambda/gokit-waffle/waffle-server/service/number"
	"github.com/gorilla/mux"

	"golang.org/x/net/context"
)

func NewNumberRouter(ctx context.Context) *mux.Router {
	numRepo := number.NewNumberRepository()
	numSvc := number.NewNumberService(numRepo)
	numRoute := mux.NewRouter()

	numInsertHandler := number.NewInsertHandler(ctx, numSvc)
	numRoute.Handle("/api/v1/number/insert", numInsertHandler).Methods("POST")
	numQueryHandler := number.NewQueryHandler(ctx, numSvc)
	numRoute.Handle("/api/v1/number/query/{user}", numQueryHandler).Methods("GET")

	return numRoute
}
