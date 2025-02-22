package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func merge(ch1 <-chan int, ch2 <-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range ch1 {
			out <- v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			out <- v
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func source(sourceFunc func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- sourceFunc(i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()

	return ch
}

func main() {
	in1 := source(func(i int) int {
		return rand.Int()
	})

	in2 := source(func(i int) int {
		return i
	})

	out := merge(in1, in2)

	for val := range out {
		fmt.Println("Value:", val)
	}
}
