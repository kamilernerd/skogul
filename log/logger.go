package skogul

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/telenornms/skogul"
)

var aggregateTimeDuration = time.Duration(time.Millisecond * 100)

func NewLog() *LogBus {
	aggregator := &LogAggregator{
		aggregatorQueue:   make(chan skogul.Metric, 100),
		outQueue:          make(chan skogul.Metric, 100),
		batch:             make(map[string]skogul.Metric),
		aggregateDuration: aggregateTimeDuration,
		aggregateTime:     *time.NewTicker(aggregateTimeDuration),
	}

	logbus := &LogBus{
		aggregator:      aggregator,
		queue:           make(chan skogul.Metric, 100),
		writers:         map[string]func(msg string){},
		logWrites:       0,
		logReads:        0,
		logWritesTicker: *time.NewTicker(rwTimeDuration),
		logReadsTicker:  *time.NewTicker(rwTimeDuration),
	}

	go aggregator.aggregate()
	go logbus.RWStats()
	go logbus.Sink()
	go logbus.Flusher()

	return logbus
}

var logger = NewLog()

func Print(v ...any) {
	logger.Push(buildLogMessage(fmt.Sprint(v...)))
}

func Printf(format string, v ...any) {
	logger.Push(buildLogMessage(fmt.Sprintf(format, v...)))
}

func Println(v ...any) {
	logger.Push(buildLogMessage(fmt.Sprintln(v...)))
}

func Error(v ...any) {
	logger.Push(buildLogMessage(fmt.Sprint(v...)))
	os.Exit(1)
}

func Errorf(format string, v ...any) error {
	logger.Push(buildLogMessage(fmt.Sprintf(format, v...)))
	return errors.New(fmt.Sprintf(format, v...))
}

func Errorln(v ...any) error {
	logger.Push(buildLogMessage(fmt.Sprintln(v...)))
	return errors.New(fmt.Sprintln(v...))
}

func Fatal(v ...any) error {
	logger.Push(buildLogMessage(fmt.Sprint(v...)))
	return errors.New(fmt.Sprint(v...))
}

func Fatalf(format string, v ...any) {
	logger.Push(buildLogMessage(fmt.Sprintf(format, v...)))
	os.Exit(1)
}

func Fatalln(v ...any) {
	logger.Push(buildLogMessage(fmt.Sprintln(v...)))
	os.Exit(1)
}

func buildLogMessage(detail string) skogul.Metric {
	funcName, _ := callerFunc()
	time := time.Now()
	container := skogul.Metric{
		Time: &time,
		Metadata: map[string]interface{}{
			"id":               funcName,
			"aggregated_count": 1,
		},
		Data: map[string]interface{}{
			"detail": detail,
		},
	}
	return container
}

func callerFunc() (string, string) {
	pc, _, _, _ := runtime.Caller(3)
	callerFuncStackPath := runtime.FuncForPC(pc).Name()
	callerFunc := strings.Split(callerFuncStackPath, "/")
	funcName := callerFunc[len(callerFunc)-1]
	return funcName, callerFuncStackPath
}

func GetLogger() *LogBus {
	return logger
}
