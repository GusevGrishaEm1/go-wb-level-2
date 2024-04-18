package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		go func() {
			wg.Wait()
			close(c)
		}()
		var ok1 bool = true
		var ok2 bool = true
		for {
			select {
			case v, ok := <-a:
				if !ok1 {
					continue
				}
				if !ok {
					wg.Done()
					ok1 = false
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok2 {
					continue
				}
				if !ok {
					wg.Done()
					ok2 = false
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
