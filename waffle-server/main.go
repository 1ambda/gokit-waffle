package main

import (
	"net/http"
	"os"
	"time"

	"github.com/1ambda/gokit-waffle/waffle-server/endpoint"
	"github.com/1ambda/gokit-waffle/waffle-server/service"
	"github.com/1ambda/gokit-waffle/waffle-server/transport"

	arg "github.com/alexflint/go-arg"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

// Build Flags
var (
	Version   = "undefined"
	BuildTime = "undefined"
	GitHash   = "undefined"
	Started   = time.Now()
)

var Environment struct {
	Mode string `arg:"env"`
	Port string `arg:"env"`
}

func main() {
	// Setup default values
	Environment.Mode = "LOCAL" // [LOCAL, DEV, PROD]
	Environment.Port = "8080"
	arg.MustParse(&Environment)

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.NewContext(logger).With(
		"ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller,
	)

	logger.Log(
		"version", Version,
		"gitHash", GitHash,
		"buildTime", BuildTime,
		"started", Started.UTC().Format(time.RFC3339),
		"mode", Environment.Mode,
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
	logger.Log("error", http.ListenAndServe(":"+Environment.Port, generateHandler))
}
