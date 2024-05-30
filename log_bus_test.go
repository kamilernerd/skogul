package skogul_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/telenornms/skogul"
)

func TestLogBusQueue(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	log_bus := skogul.NewLogBus()

	go func() {
		i := 0
		for {
			log_bus.Push(skogul.LogContainer{
				Id: time.Now().String(),
				Time: time.Now(),
				Detail: fmt.Sprintf("Log bus push %d", i),
			})
			i++
		}
	}()

	cap := make(chan int, 1)

	go log_bus.Sink()

	time.Sleep(10 * time.Second)

	cap <- 1
}
