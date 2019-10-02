package main

import "fmt"

type Car interface {
	Go()
	Stop()
}

type Lambo struct {
	Model string
}

func (l *Lambo) Go() {
	fmt.Println("Lambo model", l.Model, "is on the way to Yogyakarta")
}

func (l *Lambo) Stop() {
	fmt.Println("Lambo model", l.Model, "has arrived at Yogyakarta")
}

func NewModel(model string) Car {
	return &Lambo{model}
}

func emptyInterface(param interface{}) {
	fmt.Println(param)
}

func main() {
	c := Lambo{"S10"}
	c.Go()
	c1 := NewModel("S11")
	c1.Go()
	c1.Stop()
	emptyInterface("Hello")
	mymap := make(map[string]interface{})
	mymap["Name"] = "Lalu Kismara Hadi"
	mymap["number"] = 7
	fmt.Println(mymap)
}
