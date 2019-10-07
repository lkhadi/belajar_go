package main

import (
	"fmt"
	"os"
)

func Select(c chan int, quit chan struct{}) {
	for {
		select {
		case <-c:
			fmt.Println("ok")
		case <-quit:
			fmt.Println("Keluar")
			os.Exit(0)
		}
	}
}

func main() {
	c := make(chan int, 2)
	quit := make(chan struct{})
	go func() {
		c <- 1
		c <- 2
		quit <- struct{}{}
	}()
	go Select(c, quit)
	select {}
}
