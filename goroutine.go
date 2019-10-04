//go routine -> run in the background
package main

import (
	"fmt"
	"time"
)

func first() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("First")
	}
}

func second() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Second")
	}
}

func main() {
	go first()
	go second()
	fmt.Println("Start")
	select {}
}
