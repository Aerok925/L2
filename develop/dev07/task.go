package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	retChan := make(chan interface{}, len(channels))

	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			retChan <- <-channel
		}(channel)
	}
	return retChan
}

func main() {
	sig := func(tim time.Duration) <-chan interface{} {
		retCahn := make(chan interface{}, 1)
		go func() {
			defer close(retCahn)
			time.Sleep(tim)
			retCahn <- 1
		}()
		return retCahn
	}

	start := time.Now()
	fmt.Println(<-or(
		sig(time.Minute*1),
		sig(time.Second*3),
	))
	fmt.Println(time.Since(start).Milliseconds())
}
