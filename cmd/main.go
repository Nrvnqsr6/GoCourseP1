package main

import (
	"context"
	"sync"
)

type iProduct interface {
	Update()
}

type iUser interface {
	Update()
}

type iOrder interface {
	Update()
}

type iCart interface {
	Update()
}

type iUnifiedStruct interface {
	GetData()
}

type UnifiedStruct struct {
	out chan int
}

func (unifiedStruct UnifiedStruct) GetData() {

}

type iStatChannel interface {
	GetData()
	GetChannel()
}

type StatChannel struct {
	in chan int
}

func (statChannel StatChannel) GetData() {

}

func (statChannel StatChannel) GetChannel() chan int {
	return statChannel.in
}

func main() {
	print("aaaaa")
	ctx := context.Background()
	ch1 := make(chan iStatChannel)
	ch2 := make(chan iStatChannel)

	MergeData(ctx, ch1, ch2)

}

func MergeData(ctx context.Context, chans ...chan iStatChannel) chan iUnifiedStruct {
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
	print(outUnifiedStruct)
	return outUnifiedStruct
}

func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	// Запускаем send goroutine
	// для каждого входящего канала в cs.
	// send копирует значения из c в out
	// до тех пор пока c не закрыт, затем вызываем wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	// Запускаем goroutine чтобы закрыть out
	// когда все send goroutine выполнены
	// Это должно начаться после вызова wg.Add.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// func main() {
// 	ch := make(chan int)
// 	out := make(chan int)
// 	go producer(ch, 100*time.Millisecond)
// 	go producer(ch, 250*time.Millisecond)
// 	go reader(out)

// 	go func() {
// 		for i := range ch {
// 			out <- i
// 		}
// 	}()
// }

// func producer(ch chan int, d time.Duration) {
//     var i int
//     for {
//         ch <- i
//         i++
//         time.Sleep(d)
//     }
// }

// func reader(out chan int) {
//     for x := range out {
//         fmt.Println(x)
//     }
// }
