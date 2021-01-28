package main

import (
	"fmt"
	"time"
)

func main() {
	chStop := make(chan int, 1)
	timerfunc(chStop)

	time.Sleep(time.Second * 5)
	chStop <- 0

	close(chStop)

	time.Sleep(time.Second * 1)

	fmt.Println("Application End.")
}

func timerfunc(stopTimer chan int) {
	go func() {
		ticker := time.NewTicker(1 * time.Second) //1 sec duration Timer

	LOOP:
		for {
			select {
			case <-ticker.C:
				fmt.Printf("now -> %v\n", time.Now())
			case <-stopTimer:
				fmt.Println("Timer stop.")
				ticker.Stop()
				break LOOP
			}
		}
		fmt.Println("timerfunc end.")
	}()
}
