//Lambda function adalah fungsi tanpa nama yang digunakan ketika membutuhkan fungsi yang dipakai hanya sekali saja untuk pemakain cepat
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Wait Group -> add,done,wait
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("first")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Second")
		wg.Done()
	}()
	fmt.Println("start")
	wg.Wait()
}
