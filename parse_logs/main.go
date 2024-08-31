package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/Go-Giter/random_practice/parse_logs/stats"
)

// 2023-12-11T12:18:29.037849Z 79.212.135.122:22765 4.623 9283 500 POST https://example.com:443/webapp/process HTTP/1.1

const (
	mainTemplate  = "Number of requests: %d\nAverage request time: %.2fsec\nTotal duration: %s\n\n"
	pathsTemplate = "%s\t%.2f\t%s\n"
)

func main() {
	log := flag.String("log.file", "./log.txt", "log file we wish to parse")
	numPathsPrint := flag.Int("num.paths", 5, "number of paths to print")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	logFile, err := os.Open(*log)
	if err != nil {
		logger.With("error", err).Error("error opening log file")
		os.Exit(1)
	}

	s := stats.New(logFile, logger)

	d, err := float2Duration(s.TotalTime)
	if err != nil {
		logger.With("error", err).Error("error parsing float")

		d = time.Nanosecond
	}

	fmt.Printf(mainTemplate,
		s.TotalLines,
		s.AverageTime,
		d.Truncate(time.Second).String())

	for i := 0; i < *numPathsPrint; i++ {
		fmt.Printf(pathsTemplate,
			s.PathStats[i].Method,
			s.PathStats[i].AverageTime,
			s.PathStats[i].Path)
	}
}

func float2Duration(f float64) (time.Duration, error) {
	d, err := time.ParseDuration(strconv.FormatFloat(f, 'f', -1, 64) + "s")
	if err != nil {
		err = fmt.Errorf("error parsing duration : %w", err)
	}

	return d, err
}
