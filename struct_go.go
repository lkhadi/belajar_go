package main

import "fmt"

type Car struct {
	Merek string
	Tahun string
	Harga string
}

func (c2 Car) Get() {
	fmt.Println(c2)
}
func main() {
	var c Car
	c = Car{
		Merek: "oke",
		Tahun: "2019",
		Harga: "Rp 190000",
	}
	c.Get()
	c1 := Car{
		Merek: "oce",
		Tahun: "2020",
		Harga: "Rp 330000",
	}
	c1.Get()
	c3 := Car{"Cevrolet", "2011", "Rp 99999"}
	c3.Get()
}
