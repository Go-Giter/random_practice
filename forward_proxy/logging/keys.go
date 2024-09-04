package logging

const (
	// LoggingErrorKey is used when logging an error from slog, consistent keys matter.
	LoggingErrorKey = "error"

	// LoggingProxiedHostKey is used to log which host we are proxying.
	LoggingProxiedHostKey = "proxied_host"
)
