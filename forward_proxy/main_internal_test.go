package main

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"testing"

	"github.com/reedobrien/checkers"
)

type safeBuf struct {
	mu sync.Mutex
	b  *bytes.Buffer
}

func (sb *safeBuf) Write(p []byte) (int, error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	return sb.b.Write(p)
}

func (sb *safeBuf) Bytes() []byte {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	return sb.b.Bytes()
}

func TestRunServer(t *testing.T) {
	t.Parallel()

	testLstnr, err := net.Listen("tcp", ":0")
	checkers.OK(t, err)

	logBucket := &safeBuf{
		b: bytes.NewBuffer(nil),
	}

	logger := slog.New(slog.NewJSONHandler(logBucket, &slog.HandlerOptions{}))

	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	th := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`OK`))
	})

	wg.Add(1)

	go runServer(testLstnr, th, wg, ctx, logger)

	testClient := &http.Client{}

	resp, err := testClient.Get("http://" + testLstnr.Addr().String())
	checkers.OK(t, err)

	defer resp.Body.Close()

	checkers.Equals(t, resp.StatusCode, http.StatusOK)

	b, err := io.ReadAll(resp.Body)
	checkers.OK(t, err)

	checkers.Equals(t, string(b), `OK`)

	cancel()

	wg.Wait()
}
