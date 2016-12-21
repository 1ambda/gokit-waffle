package main

import (
	"net/http"
	"os"

	"github.com/1ambda/gokit-waffle/waffle-server/config"
	"github.com/1ambda/gokit-waffle/waffle-server/service"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

func main() {

	flag := config.GetFlag()
	env := config.GetEnvironment()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.NewContext(logger).With(
		"ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller,
	)

	logger.Log(
		"version", flag.Version,
		"gitHash", flag.GitHash,
		"buildTime", flag.BuildTime,
		"started", flag.Started,
		"mode", env.Mode,
		"port", env.Port,
	)

	// Start
	ctx := context.Background()

	numRepo := service.NewNumberRepository()
	numHandler := service.NewInsertHandler(ctx, numRepo)
	numRoute := mux.NewRouter()
	numRoute.Handle("/api/v1/number/insert", numHandler).Methods("POST")

	mux := http.NewServeMux()
	mux.Handle("/api/v1/number/", numRoute)

	http.Handle("/", mux)

	logger.Log("error", http.ListenAndServe(env.Port, nil))
}
