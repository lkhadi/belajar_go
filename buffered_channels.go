package main

import "fmt"

type Car struct {
	Model string
}

func main() {
	c := make(chan *Car, 3)
	go func() {
		c <- &Car{"Lambo"}
		c <- &Car{"BMW"}
		c <- &Car{"Ford"}
		c <- &Car{"Tesla"}
		close(c)
	}()
	for i := range c {
		fmt.Println(i.Model)
	}
}
