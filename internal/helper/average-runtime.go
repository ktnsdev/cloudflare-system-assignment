package helper

import (
	"strings"
	"time"
)

var averageEncodingRuntime = time.Millisecond * 0
var averageDecodingRuntime = time.Millisecond * 0
var totalEncodingRuntime = time.Millisecond * 0
var totalDecodingRuntime = time.Millisecond * 0
var encodingTimesCalled = 0
var decodingTimesCalled = 0

type Runtimes struct {
	Path      string
	Timestamp time.Time
	Runtime   time.Duration
}

var allRuntimes []Runtimes

func GetAverageEncodingRuntime() time.Duration {
	return averageEncodingRuntime
}

func GetAverageDecodingRuntime() time.Duration {
	return averageDecodingRuntime
}

func GetAllRuntimes() []Runtimes {
	return allRuntimes
}

func addToAverageRuntime(path string, runtime time.Duration, timestamp time.Time) {
	allRuntimes = append(allRuntimes, Runtimes{path, timestamp, runtime})

	if strings.HasPrefix(path, "/auth") {
		encodingTimesCalled++
		totalEncodingRuntime += runtime
		averageEncodingRuntime = totalEncodingRuntime / time.Duration(encodingTimesCalled)
	} else if strings.HasPrefix(path, "middleware") || strings.HasPrefix(path, "/verify") {
		decodingTimesCalled++
		totalDecodingRuntime += runtime
		averageDecodingRuntime = totalDecodingRuntime / time.Duration(decodingTimesCalled)
	}
}

func CalculateElapsedTime(path string) func() {
	timeStart := time.Now()
	return func() {
		addToAverageRuntime(path, time.Since(timeStart), timeStart)
	}
}
