package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{} = func(channels ...<-chan interface{}) <-chan interface{} {
		mainChannel := make(chan interface{}, len(channels))
		go func() {
			defer close(mainChannel)
			wg := sync.WaitGroup{}
			for _, ch := range channels {
				wg.Add(1)
				go func(ch <-chan interface{}) {
					defer wg.Done()
					for v := range ch {
						fmt.Print("Received: ", v)
						mainChannel <- v
					}
				}(ch)
			}
			wg.Wait()
		}()
		return mainChannel
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
			c <- struct{}{}
		}()
		return c
	}

	start := time.Now()
	for range or(
		sig(2*time.Minute),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(30*time.Second),
		sig(1*time.Minute),
	) {
	}

	fmt.Printf("fone after %v", time.Since(start))
}
