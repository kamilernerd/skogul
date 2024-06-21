package skogul

import (
	"fmt"
	"time"

	"github.com/telenornms/skogul"
)

type LogAggregator struct {
	aggregatorQueue   chan skogul.Metric
	outQueue          chan skogul.Metric
	batch             map[string]skogul.Metric
	aggregateTime     time.Ticker
	aggregateDuration time.Duration
}

func (ag *LogAggregator) aggregate() {
	for {
		select {
		case q := <-ag.aggregatorQueue:
			key := ag.makeMapKey(q)
			if val, ok := ag.batch[key]; ok {
				val.Metadata["aggregated_count"] = val.Metadata["aggregated_count"].(int) + 1
			} else {
				ag.batch[key] = q
			}
		case <-ag.aggregateTime.C:
			for k, v := range ag.batch {
				ag.outQueue <- v
				delete(ag.batch, k)
			}
			ag.aggregateTime.Reset(ag.aggregateDuration)
		default:
			continue
		}
	}
}

func (ag *LogAggregator) makeMapKey(q skogul.Metric) string {
	detail := q.Data["detail"].(string)
	id := q.Metadata["id"].(string)
	return fmt.Sprintf("%s-%s", id, detail)
}
