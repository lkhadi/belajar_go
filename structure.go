package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	arr := []string{"ok", "oke", "k"}
	for i, value := range arr {
		fmt.Println(i, value)
	}
	//SWITCH CASE
	number := 4
	switch number {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3, 4:
		fmt.Println("the number is 3 or 4")
	default:
		fmt.Println("default")
	}
	switch {
	case number == 1:
		fmt.Println(1)
	case number == 2:
		fmt.Println(2)
	case number == 3, number == 4:
		fmt.Println("the number is 3 or 4")
	}
	//IF STATEMENT
	status := true
	if status {
		fmt.Println("Nice")
	} else {
		fmt.Println("Bad")
	}
}
