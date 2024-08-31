package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	addr := flag.String("addr", ":8080", "where we want to listen")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	lstnr, err := net.Listen("tcp", *addr)
	if err != nil {
		logger.With("error", err).Error("error listening")
		os.Exit(1)
	}

	s := &http.Server{
		Handler: handlerFunc(logger),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.Serve(lstnr); !errors.Is(err, http.ErrServerClosed) {
			logger.With("error", err).Error("unexpected error returned")
			os.Exit(1)
		}
	}()

	<-stop

	sdCtx, sdCancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = s.Shutdown(sdCtx)

	sdCancel()
}

func handlerFunc(logger *slog.Logger) http.HandlerFunc {
	healthMetrics := promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "health_processing_seconds",
		Help: "Time spent processing /health endpoint",
	})

	rndm := rand.New(rand.NewSource(time.Now().Unix()))
	h := promhttp.Handler()

	return func(w http.ResponseWriter, r *http.Request) {
		endPoint := r.URL.RequestURI()

		switch endPoint {
		case "/health":
			start := time.Now()
			if rndm.Int()%3 == 0 {
				<-time.After(time.Second)
			}

			_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
			if err != nil {
				logger.With("error", err).Error("error writing /health response")
			}

			healthMetrics.Observe(float64(time.Since(start).Seconds()))

			return
		case "/metrics":
			h.ServeHTTP(w, r)

			return
		default:
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(`HTTP/1.0 404 Not Found`))
			if err != nil {
				logger.With("error", err).Error("error writing 404 response")
			}

			return
		}
	}
}
