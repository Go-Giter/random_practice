package proxyhandler

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"scratch/random_practice/forward_proxy/config"
	"scratch/random_practice/forward_proxy/logging"
	"sync"
)

type ProxyHandler struct {
	httpClient *http.Client
	sp         *sync.Pool
	logger     *slog.Logger
	cfg        config.Config
}

// ErrEmptyHost is logged when we receive a request with an empty host header ( can this even happen?)
var ErrEmptyHost = errors.New("request received with an empty host header")

func New(httpClient *http.Client, cfg config.Config, logger *slog.Logger) *ProxyHandler {
	return &ProxyHandler{
		logger:     logger,
		httpClient: httpClient,
		cfg:        cfg,
		sp: &sync.Pool{
			New: func() any {
				return bytes.NewBuffer(nil)
			},
		},
	}
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h := r.Host
	s := r.URL.Scheme

	if h == "" {
		ph.logger.With(logging.LoggingErrorKey, ErrEmptyHost).Error("invalid request received")

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	ph.logger.With(logging.LoggingProxiedHostKey, h).Debug("handling proxy request")

	req, err := http.NewRequestWithContext(ctx, r.Method, s+"://"+h, nil)
	if err != nil {
		ph.logger.With(logging.LoggingErrorKey, err).Error("error creating request")

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	shost, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ph.logger.With(logging.LoggingErrorKey, err).Error("error splitting RemoteAddr")

		shost = ""
	}

	req.Header.Add("X-Forwarded-For", shost)

	req.Header.Add("proxy", "dvt")

	resp, err := ph.httpClient.Do(req)
	if err != nil {
		ph.logger.With(logging.LoggingErrorKey, err).Error("error executing proxied request")

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	defer resp.Body.Close()

	buf := ph.sp.Get().(*bytes.Buffer)

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		ph.logger.With(logging.LoggingErrorKey, err).Error("error copying resp.Body")

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	w.WriteHeader(resp.StatusCode)

	_, err = w.Write(buf.Bytes())
	if err != nil {
		ph.logger.With(logging.LoggingErrorKey, err).Error("error writing response")
	}

	buf.Reset()

	ph.sp.Put(buf)
}
