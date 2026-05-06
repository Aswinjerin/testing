package main

import (
	"fmt"
	"sync"
)

func main() {
	var a int

	fmt.Print("Enter the number ")
	fmt.Scan(&a)
	val := a / 10
	var start int
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		end := val * i
		wg.Add(1)
		go func(val int, l, k int) {
			defer wg.Done()
			x := 0
			for j := l; j <= k; j++ {
				x = j
			}
			fmt.Printf("%d val %d \n", val, x)
		}(i, start, end)
		start = end
	}
	wg.Wait()

}
