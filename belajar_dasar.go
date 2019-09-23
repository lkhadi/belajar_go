package main

import "fmt"

func main() {
	// declaration()
	// _Array()
	// number1, number2 := 1, 2
	// _Pointer(&number1, &number2)
	basicPointer()
}

func basicPointer() {
	m1 := 2
	ptr := &m1
	*ptr = 9
	fmt.Println(m1)
}

func _Pointer(m1, m2 *int) {
	var tmp int
	tmp = *m1
	*m1 = *m2
	*m2 = tmp
}

func declaration() {
	var number int
	number = 21
	number2 := 22
	var (
		number3   = 23
		_sentence = "alskdfj"
	)
	fmt.Println(number+number2+number3, _sentence)
}

func _Array() {
	var number []int
	number = []int{1, 2, 3, 4}
	number2 := []int{5, 6}
	number2 = append(number2, 7)
	var (
		number3   = []int{8, 9}
		_sentence = []string{"nice", "ok"}
	)
	fmt.Println(number, number2, number3, _sentence)
}
