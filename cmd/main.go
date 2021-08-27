package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caiowWillian/partner-service/cmd/route"
	"github.com/caiowWillian/partner-service/internal/partner"
	"github.com/caiowWillian/partner-service/pkg/mongo"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

const port = 5555

func main() {
	jaegerCfg := config.Configuration{
		ServiceName: "partner-service",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, closer, err := jaegerCfg.NewTracer()

	if err != nil {
		os.Exit(0)
	}

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	ctx := context.Background()

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		fmt.Println("==============")
		fmt.Println(os.Getenv("mongo_uri"))
		configureMongoDb()
		address := fmt.Sprintf(":%d", port)
		r := mux.NewRouter()
		r.Handle("/metrics", promhttp.Handler())
		route.MakeRoutes(ctx, r)

		fmt.Println("listening on port", address)
		errs <- http.ListenAndServe(address, r)
	}()

	level.Error(logger).Log("exit", <-errs)
}

func configureMongoDb() {
	err := partner.NewRepository(mongo.NewMongo()).CreateIndexes()
	fmt.Println(err)
}
