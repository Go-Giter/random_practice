package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"scratch/random_practice/forward_proxy/logging"
	proxyhandler "scratch/random_practice/forward_proxy/proxy_handler"
	"sync"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultClientTimeout = time.Second * 5

	defaultSeverStopGrace = time.Second * 10
)

func main() {
	proxyPort := flag.String("listen.port", "8989", "port we should listen on")
	metricsPort := flag.String("metrics.port", "8080", "port for prom to listen on")
	verbose := flag.Bool("verbose", false, "enable debug logs?")
	flag.Parse()

	slogOpts := &slog.HandlerOptions{}
	if *verbose {
		slogOpts.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, slogOpts))

	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	proxyLstnr, err := net.Listen("tcp", ":"+*proxyPort)
	if err != nil {
		logger.With(logging.LoggingErrorKey, err).Error("unable to listen on " + *proxyPort)

		os.Exit(1)
	}

	c := &http.Client{
		Timeout: defaultClientTimeout,
	}

	ph := proxyhandler.New(c, logger)

	wg.Add(1)

	go runServer(proxyLstnr, ph, wg, ctx, logger)

	metricsLstnr, err := net.Listen("tcp", ":"+*metricsPort)
	if err != nil {
		logger.With(logging.LoggingErrorKey, err).Error("unable to listen on " + *metricsPort)

		os.Exit(1)
	}

	wg.Add(1)

	go runMetricsServer(metricsLstnr, wg, ctx, logger)

	<-stop
	cancel()

	wg.Wait()
}

func runServer(proxyLstnr net.Listener,
	ph http.Handler,
	wg *sync.WaitGroup,
	ctx context.Context,
	logger *slog.Logger,
) {
	defer wg.Done()

	s := &http.Server{
		Handler: ph,
	}

	go func() {
		if err := s.Serve(proxyLstnr); !errors.Is(err, http.ErrServerClosed) {
			logger.With(logging.LoggingErrorKey, err).Error("error starting proxy server")

			os.Exit(1)
		}
	}()

	<-ctx.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), defaultSeverStopGrace)

	if err := s.Shutdown(stopCtx); err != nil {
		log.Println("error shutting down: ", err)
	}

	stopCancel()

	logger.Info("proxy server exiting")
}

func runMetricsServer(metricsLstnr net.Listener, wg *sync.WaitGroup, ctx context.Context, logger *slog.Logger) {
	defer wg.Done()

	s := &http.Server{
		Handler: promhttp.Handler(),
	}

	go func() {
		if err := s.Serve(metricsLstnr); !errors.Is(err, http.ErrServerClosed) {
			logger.With(logging.LoggingErrorKey, err).Error("error starting metrics server")

			os.Exit(1)
		}
	}()

	<-ctx.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), defaultSeverStopGrace)

	if err := s.Shutdown(stopCtx); err != nil {
		log.Println("error shutting down: ", err)
	}

	stopCancel()

	logger.Info("metrics server exiting")
}
