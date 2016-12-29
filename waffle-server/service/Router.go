package service

import (
	"github.com/1ambda/gokit-waffle/waffle-server/service/number"
	"github.com/1ambda/gokit-waffle/waffle-server/service/user"
	"github.com/gorilla/mux"

	"golang.org/x/net/context"
)

func NewNumberRouter(ctx context.Context, r number.NumberRepository) *mux.Router {
	svc := number.NewNumberService(r)
	route := mux.NewRouter().StrictSlash(false)

	insertHandler := number.NewInsertHandler(ctx, svc)
	route.Handle("/api/v1/number/insert", insertHandler).Methods("POST")

	return route
}

func NewUserRouter(ctx context.Context, r number.NumberRepository) *mux.Router {
	svc := user.NewUserService(r)
	route := mux.NewRouter()

	usersHandler := user.NewUserListHandler(ctx, svc)
	route.Handle("/api/v1/user", usersHandler).Methods("GET")
	userHandler := user.NewUserHandler(ctx, svc)
	route.Handle("/api/v1/user/{user}", userHandler).Methods("GET")

	return route
}
