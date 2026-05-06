package demo

import (
	"fmt"
)

func cha() {
	//const total = 1000
	//with channel

	ch := make(chan int)

	for i := 0; i < 1000; i++ {
		go func() {
			ch <- 1
		}()
	}

	count := 0

	for i := 0; i < 1000; i++ {
		count += <-ch
	}

	//without channel
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	fmt.Println("Final count without channel:", counter)
	fmt.Println("Final count:", count)

}
