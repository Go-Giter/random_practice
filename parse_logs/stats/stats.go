package stats

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type (
	PathStats struct {
		Path        string
		Method      string
		AverageTime float64
		TotalCount  int
		TotalTime   float64
	}

	Stats struct {
		PathStats   []PathStats
		TotalLines  int
		TotalTime   float64
		AverageTime float64
	}
)

func New(logFile *os.File, logger *slog.Logger) *Stats {
	s := &Stats{
		TotalLines: 0,
		TotalTime:  0.0,
		PathStats:  []PathStats{},
	}

	s.parseLog(logFile, logger)

	return s
}

func (s *Stats) parseLog(logFile *os.File, logger *slog.Logger) {
	scanner := bufio.NewScanner(logFile)
	uniquePathStats := make(map[string]map[string]PathStats)

	for scanner.Scan() {
		line := strings.ReplaceAll(
			strings.ReplaceAll(scanner.Text(), `"`, ""), "https://example.com:443", "")

		logFields := strings.Fields(line)

		f, err := strconv.ParseFloat(logFields[2], 64)
		if err != nil {
			slog.With("error", err).Error("error parsing float")

			f = 0.0
		}

		s.TotalTime += f

		s.TotalLines++

		s.parsePathStats(logFields, uniquePathStats, logger)
	}

	s.AverageTime = s.TotalTime / float64(s.TotalLines)
	s.calcAvgAndSort(uniquePathStats)
}

func (s Stats) parsePathStats(logFields []string,
	uniquePathStats map[string]map[string]PathStats,
	logger *slog.Logger,
) {
	time := logFields[2]
	method := logFields[5]
	path := logFields[6]

	f, err := strconv.ParseFloat(time, 64)
	if err != nil {
		logger.With("error", err).Error("error parsing time")

		f = 0.0
	}

	_, pok := uniquePathStats[path]
	if !pok {
		uniquePathStats[path] = make(map[string]PathStats)

		_, mok := uniquePathStats[path][method]
		if !mok {
			uniquePathStats[path][method] = PathStats{
				Path:       path,
				TotalCount: 1,
				TotalTime:  f,
				Method:     method,
			}
		}

		return
	}

	v := uniquePathStats[path][method]
	v.TotalCount++
	v.TotalTime += f
	uniquePathStats[path][method] = v
}

func (s *Stats) calcAvgAndSort(uniquePathStats map[string]map[string]PathStats) {
	for pathk := range uniquePathStats {
		for methodk := range uniquePathStats[pathk] {
			ps := uniquePathStats[pathk][methodk]

			ps.AverageTime = ps.TotalTime / float64(ps.TotalCount)

			s.PathStats = append(s.PathStats, ps)
		}
	}

	for i := 0; i < len(s.PathStats)-1; i++ {
		for j := 0; j < len(s.PathStats)-i-1; j++ {
			if s.PathStats[j].AverageTime < s.PathStats[j+1].AverageTime {
				s.PathStats[j+1], s.PathStats[j] = s.PathStats[j], s.PathStats[j+1]
			}
		}
	}
}
