package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type iUpdatable interface {
	Update() uuid.UUID
}

type TypeWithUpdate struct {
	id uuid.UUID
}

func (update TypeWithUpdate) Update() uuid.UUID {
	return update.id
}

type iUnifiedStruct interface {
	GetData() uuid.UUID
}

type UnifiedStruct struct {
	model iUpdatable
}

func (unifiedStruct UnifiedStruct) GetData() uuid.UUID {
	return unifiedStruct.model.Update()
}

type iStatisticChannel interface {
	GetData() uuid.UUID
}

type StatisticChannel struct {
	model iUpdatable
}

func (statChannel StatisticChannel) GetData() uuid.UUID {
	return statChannel.model.Update()
}

func main() {
	ctx := context.Background()
	ch1 := make(chan iStatisticChannel)
	ch2 := make(chan iStatisticChannel)
	go func() {
		ch1 <- StatisticChannel{model: TypeWithUpdate{id: uuid.New()}}
		ch2 <- StatisticChannel{model: TypeWithUpdate{id: uuid.New()}}
	}()

	MergeDataInterfaceArg(ctx, StatisticChannel{model: TypeWithUpdate{id: uuid.New()}}, StatisticChannel{model: TypeWithUpdate{id: uuid.New()}})
	a := MergeDataChanArg(ctx, ch1, ch2)
	a1 := <-a
	fmt.Println(a1.GetData())
	a2 := <-a
	fmt.Println(a2.GetData())
}

func MergeDataChanArg(ctx context.Context, chans ...chan iStatisticChannel) chan iUnifiedStruct {
	var wg sync.WaitGroup

	outUnifiedStruct := make(chan iUnifiedStruct)

	wg.Add(len(chans))

	for _, c := range chans {
		go func(c chan iStatisticChannel) {
			defer wg.Done()
			for val := range c {
				outUnifiedStruct <- val
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(outUnifiedStruct)
	}()

	return outUnifiedStruct
}

func MergeDataInterfaceArg(ctx context.Context, chans ...iStatisticChannel) chan iUnifiedStruct {
	var wg sync.WaitGroup

	outUnifiedStruct := make(chan iUnifiedStruct)

	wg.Add(len(chans))

	for _, c := range chans {
		go func(c iStatisticChannel) {
			defer wg.Done()
			outUnifiedStruct <- c
		}(c)
	}

	go func() {
		wg.Wait()
		close(outUnifiedStruct)
	}()

	return outUnifiedStruct
}
