package skogul

import (
	// "encoding/json"
	"fmt"
	"time"
)

type LogBus struct {
	queue           chan Metric
	aggregator      *LogAggregator
	writers         map[string]func(msg string)
	logWrites       int
	logReads        int
	logWritesTicker time.Ticker
	logReadsTicker  time.Ticker
}

var rwTimeDuration = time.Duration(time.Second * 1)

func (bus *LogBus) RWStats() {
	for {
		select {
		case <-bus.logWritesTicker.C:
			bus.logWrites = 0
			bus.logWritesTicker.Reset(rwTimeDuration)
		case <-bus.logReadsTicker.C:
			bus.logReads = 0
			bus.logReadsTicker.Reset(rwTimeDuration)
		default:
			continue
		}
	}
}

func (bus *LogBus) Push(v Metric) {
	select {
	case bus.queue <- v:
		bus.logWrites++
	default:
		<-bus.queue
	}
}

func (bus *LogBus) Sink() {
	for {
		select {
		case q := <-bus.queue:
			bus.logReads++
			bus.aggregator.aggregatorQueue <- q
		default:
			continue
		}
	}
}

var LogHandler HandlerRef

func (bus *LogBus) Flusher() {
	for {
		select {
		case q := <-bus.aggregator.outQueue:
			metrics := make([]*Metric, 1)
			metrics[0] = &q
			c := Container{
				Metrics: metrics,
			}

			err := LogHandler.H.Send(&c)
			if err != nil {
				fmt.Println(err)
				break
			}
		default:
			continue
		}
	}
}
