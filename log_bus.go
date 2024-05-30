package skogul

import (
	// "flag"
	"fmt"
	"sync"
	"time"
)

type LogBus struct {
	queue           chan LogContainer
	out             chan LogContainer
	sidequeue       chan LogContainer
	writers         map[string]func(msg string)
	logWrites       int
	logReads        int
	logWritesTicker time.Ticker
	logReadsTicker  time.Ticker
	logBatch        map[string][]LogContainer
	mutex           sync.RWMutex
}

var rwTimeDuration = time.Duration(time.Second * 1)

func NewLogBus() *LogBus {
	// silentFlag := flag.Bool("sout", false, "Silence stdout output")

	logbus := &LogBus{
		queue:           make(chan LogContainer, 100),
		sidequeue:       make(chan LogContainer, 100),
		out:             make(chan LogContainer, 100),
		writers:         map[string]func(msg string){},
		logWrites:       0,
		logReads:        0,
		logWritesTicker: *time.NewTicker(rwTimeDuration),
		logReadsTicker:  *time.NewTicker(rwTimeDuration),
		logBatch:        make(map[string][]LogContainer),
	}

	logbus.AddLogWriter("default", func(msg string) {
		// if !*silentFlag {
			fmt.Print(msg)
		// }
	})

	return logbus
}

// RWStats check whenever write ticker or read ticker reaches timeout
// resets both timers and their coresponding counter
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

// Pushes arg v to the queue and increases write counter
// When queue is full the log message is being pushed to sidequeue
// which will skip any aggregation steps and procedes to log the message
func (bus *LogBus) Push(v LogContainer) {
	select {
	case bus.queue <- v:
		bus.logWrites++
	default:
		bus.sidequeue <- <-bus.queue
	}
}

// Listens to the main queue and reads messages of it
// Aggregates duplicated messages and increases reads counter
// Proceeds to push message to the out queue
func (bus *LogBus) Sink() {
	for {
		select {
		case q := <-bus.queue:
			bus.logReads++
			bus.mutex.Lock()
			bus.logBatch[q.Detail] = append(bus.logBatch[q.Detail], q)
			bus.mutex.Unlock()
			bus.out <- q
		default:
			continue
		}
	}
}

// Reads messages of the out queue and builds a final log string
// proceeds to write the log message to the registered writers
// while cleaning up the aggregates.
// Also listens for the sidequeue in case main queue is full and
// builds the log message and then writes to the registered writers directtly
func (bus *LogBus) Flusher() {
	for {
		select {
		case <-bus.out:
			bus.mutex.RLock()
			for k, v := range bus.logBatch {
				if len(v) > 0 {
					log := v[0]

					msg := ""
					if len(v) > 1 {
						msg = fmt.Sprintf("%s %s %s - %d duplicate messages\n", log.Id, log.Time, log.Detail, len(v))
					} else {
						msg = fmt.Sprintf("%s %s %s\n", log.Id, log.Time, log.Detail)
					}
					for _, k := range bus.writers {
						k(msg)
					}
					delete(bus.logBatch, k)
				}
			}
			bus.mutex.RUnlock()
		case l := <-bus.sidequeue:
			for _, k := range bus.writers {
				k(fmt.Sprintf("%s %s %s\n", l.Id, l.Time, l.Detail))
			}
		}
	}
}

func (bus *LogBus) AddLogWriter(name string, cb func(msg string)) {
	if _, ok := bus.writers[name]; !ok {
		bus.writers[name] = cb
		return
	}
	panic("log writer with this name is already defined")
}
