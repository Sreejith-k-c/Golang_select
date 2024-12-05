package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("final")
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	go func() {
		ch1 <- "msg from ch1"
	}()
	go func() {
		ch2 <- "msg from ch2"
	}()
	// time.Sleep(time.Second * 2)
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case msg := <-ch1:
				fmt.Println(msg)
			case msg := <-ch2:
				fmt.Println(msg)
			case msg := <-ticker.C:
				fmt.Println("ticker", msg)
			case <-done:
				fmt.Println("done")
			}
		}
	}()
	time.Sleep(time.Second * 5)

	done <- true
}
