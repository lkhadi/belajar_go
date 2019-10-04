//unbuffered channel memiliki ukuran 1 yang berarti hanya dapat menampung 1 value
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	val := <-c
	fmt.Println(val)
	go func() {
		c <- 2
	}()
	val = <-c
	time.Sleep(time.Second * 1)
	fmt.Println(val)
	fmt.Println(c)
}
