package service

import (
	"context"
	"part1/internal/model"
	"sync"
)

func MergeData(ctx context.Context, chans ...chan model.Status) chan model.Status {
	outUnifiedChan := make(chan model.Status)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, channel := range chans {
		go func(channel <-chan model.Status) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case status := <-channel:
					outUnifiedChan <- status
				}
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(outUnifiedChan)
	}()

	return outUnifiedChan
}

// type Merger struct {
// 	output chan IMergeable
// 	wg     *sync.WaitGroup
// }

// func CreateMerger(ctx context.Context, chans ...chan IMergeable) Merger {
// 	merger := Merger{wg: new(sync.WaitGroup), output: make(chan IMergeable)}

// }

// func (m *Merger) AddChannels(ctx context.Context, chans ...chan int) {
// 	for _, ch := range chans {
// 		m.AddChannel(ctx, ch)
// 	}
// }

// func (m *Merger) AddChannel(ctx context.Context, ch chan int) {
// 	m.wg.Add(1)
// 	go m.ListenChannel(ctx, m.wg, ch, m.output)
// }

// func (merger *Merger) ListenChannel(ctx context.Context, chans ...chan IMergeable) chan IMergeable {

// 	for _, channel := range chans {
// 		go func(channel chan IMergeable) {
// 			defer wg.Done()

// 			for {
// 				select {
// 				case <-ctx.Done():
// 					return
// 				case msg := <-channel:
// 					outUnifiedChan <- msg
// 				}
// 			}
// 		}(channel)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(outUnifiedChan)
// 	}()

// 	return outUnifiedChan
// }
