package main

import (
	"net/http"
	"os"

	"github.com/1ambda/gokit-waffle/waffle-server/config"
	"github.com/1ambda/gokit-waffle/waffle-server/endpoint"
	"github.com/1ambda/gokit-waffle/waffle-server/service"
	"github.com/1ambda/gokit-waffle/waffle-server/transport"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
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
	svc := service.NumberServiceInst{}

	generateHandler := httptransport.NewServer(
		ctx,
		endpoint.CreateGenerationEndpoint(svc),
		transport.DecodeGenerateRequest,
		transport.EncodeGenerateResponse,
	)

	http.Handle("/api/v1/generate", generateHandler)
	logger.Log("error", http.ListenAndServe(env.Port, generateHandler))
}
