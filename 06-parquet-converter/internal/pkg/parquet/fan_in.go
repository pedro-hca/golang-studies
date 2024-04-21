package parquet

import (
	"sync"

	"github.com/apache/arrow/go/v16/arrow"
)

func FanIn(chans ...<-chan []arrow.Record) chan []arrow.Record {
	out := make(chan []arrow.Record)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))
	for _, c := range chans {
		go func(arrow <-chan []arrow.Record) {
			for rec := range arrow {
				out <- rec
			}
			wg.Done()
		}(c)
	}
	return out
}
