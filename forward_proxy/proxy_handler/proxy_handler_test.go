package proxyhandler_test

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"scratch/random_practice/forward_proxy/config"
	proxyhandler "scratch/random_practice/forward_proxy/proxy_handler"

	"github.com/reedobrien/checkers"
)

func TestServeHTTP(t *testing.T) {
	t.Parallel()

	t.Run("MethodGet", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			checkers.Equals(t, r.Header.Get("proxy"), "dvt")

			_, err := w.Write([]byte(`HELLO`))

			checkers.OK(t, err)
		}))

		tut := proxyhandler.New(ts.Client(), config.Config{}, slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{})))
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
		checkers.OK(t, err)

		rr := httptest.NewRecorder()

		tut.ServeHTTP(rr, req)

		checkers.Equals(t, rr.Code, http.StatusOK)

		b, err := io.ReadAll(rr.Body)
		checkers.OK(t, err)

		checkers.Equals(t, string(b), "HELLO")
	})
}
