package main

import (
	"net/http"
	"os"

	"github.com/1ambda/gokit-waffle/waffle-server/config"
	"github.com/1ambda/gokit-waffle/waffle-server/service"

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

	http.Handle("/api/v1/insert", service.CreateInsertHandler(ctx))
	logger.Log("error", http.ListenAndServe(env.Port, nil))
}
