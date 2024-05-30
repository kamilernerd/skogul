package skogul

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type LogContainer struct {
	Id     string    `json:"id"`
	Time   time.Time `json:"time"`
	Caller string    `json:"caller"`
	Detail string    `json:"detail"`
}

type Log struct {
	LogBus *LogBus
}

func NewLog() *Log {
	log := &Log{
		LogBus: NewLogBus(),
	}

	go log.LogBus.RWStats()
	go log.LogBus.Sink()
	go log.LogBus.Flusher()

	return log
}

var logger = NewLog()

func Print(v ...any) {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprint(v...)))
}

func Printf(format string, v ...any) {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprintf(format, v...)))
}

func Println(v ...any) {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprintln(v...)))
}

func Error(v ...any) {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprint(v...)))
	os.Exit(1)
}

func Errorf(format string, v ...any) error {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprintf(format, v...)))
	return errors.New(fmt.Sprintf(format, v...))
}

func Errorln(v ...any) error {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprintln(v...)))
	return errors.New(fmt.Sprintln(v...))
}

func Fatal(v ...any) error {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprint(v...)))
	return errors.New(fmt.Sprint(v...))
}

func Fatalf(format string, v ...any) {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprintf(format, v...)))
	os.Exit(1)
}

func Fatalln(v ...any) {
	logger.LogBus.Push(logger.buildLogMessage(fmt.Sprintln(v...)))
	os.Exit(1)
}

func (l *Log) buildLogMessage(detail string) LogContainer {
	funcName, _ := l.callerFunc()
	container := LogContainer{
		Id:     funcName,
		Time:   time.Now(),
		Detail: detail,
	}
	return container
}

func (l *Log) callerFunc() (string, string) {
	pc, _, _, _ := runtime.Caller(3)
	callerFuncStackPath := runtime.FuncForPC(pc).Name()
	callerFunc := strings.Split(callerFuncStackPath, "/")
	funcName := callerFunc[len(callerFunc)-1]
	return funcName, callerFuncStackPath
}
