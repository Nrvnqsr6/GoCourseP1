package main

import (
	"context"
	"sync"
)

type iUpdatable interface {
	Update() int
}

type Update struct {
	id int
}

func (update Update) Update() int {
	return update.id
}

type iUnifiedStruct interface {
	GetData() int
}

type UnifiedStruct struct {
	model iUpdatable
}

func (unifiedStruct UnifiedStruct) GetData() int {
	return unifiedStruct.model.Update()
}

type iStatChannel interface {
	GetData() int
}

type StatChannel struct {
	model iUpdatable
}

func (statChannel StatChannel) GetData() int {
	return statChannel.model.Update()
}

func main() {
	//print("aaaaaa")
	ctx := context.Background()
	ch1 := make(chan iStatChannel)
	ch2 := make(chan iStatChannel)
	//cl := StatChannel{model: Update{id: 10}}
	//print(cl.model)
	go func() {
		ch1 <- StatChannel{model: Update{id: 10}}
		ch2 <- StatChannel{model: Update{id: 5}}
	}()
	MergeDataChanArg(ctx, ch1, ch2)

	MergeDataInterfaceArg(ctx, StatChannel{model: Update{id: 10}}, StatChannel{model: Update{id: 5}})
	// a1 := <-a
	// print(a1.GetData())
	// a2 := <-a
	// print(a2.GetData())
}

func MergeDataChanArg(ctx context.Context, chans ...chan iStatChannel) chan iUnifiedStruct {
	//outUnifiedStruct := UnifiedStruct{out: make(chan int)}
	var wg sync.WaitGroup

	outUnifiedStruct := make(chan iUnifiedStruct)

	send := func(c chan iStatChannel) {
		for n := range c {
			outUnifiedStruct <- n
		}
		wg.Done()
	}

	wg.Add(len(chans))

	for _, c := range chans {
		go send(c)
	}
	return outUnifiedStruct
}

func MergeDataInterfaceArg(ctx context.Context, chans ...iStatChannel) chan iUnifiedStruct {
	//outUnifiedStruct := UnifiedStruct{out: make(chan int)}
	var wg sync.WaitGroup

	outUnifiedStruct := make(chan iUnifiedStruct)

	send := func(c iStatChannel) {
		outUnifiedStruct <- c
		wg.Done()
	}

	wg.Add(len(chans))

	for _, c := range chans {
		go send(c)
	}
	return outUnifiedStruct
}
