package demo

import (
	"fmt"
	"sync"
)

var (
	balance int = 0
	mu      sync.Mutex
)

func deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	balance += amount
	mu.Unlock()
}

func min() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go deposit(1, &wg)
	}
	wg.Wait()

	fmt.Printf("Final Balance: %d\n", balance)
}

// package main

// import (
// 	"fmt"
// )

// var balance int = 0

// func deposit(amount int) {

// 	balance += amount
// }
// func main() {

// 	for i := 0; i < 1000; i++ {

// 		go deposit(1)
// 	}

// 	fmt.Printf("Final Balance: %d\n", balance)
// }

// package demo

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func Great() {
// 	str := "Hello"
// 	a := 4
// 	fmt.Println(str, a)
// }
// func printnum(i int) {
// 	fmt.Println("Number", i)
// }
// func routin(a int, wg *sync.WaitGroup) {
// 	wg.Done()
// 	fmt.Println("Num", a)
// }

// func gorout() {

// 	ch := make(chan int)
// 	go func() {
// 		ch <- 5
// 	}()
// 	result := <-ch
// 	go fmt.Println("the channel", result)

// 	var wg sync.WaitGroup
// 	for i := 1; i < 5; i++ {
// 		go printnum(i)
// 	}
// 	for j := 0; j < 6; j++ {
// 		wg.Add(1)
// 		go routin(j, &wg)
// 	}
// 	wg.Wait()

// 	go Great()
// 	go fmt.Println("aswin")
// 	go fmt.Println("World")
// 	go fmt.Println("map")
// 	go fmt.Println("area")
// 	time.Sleep(1 * time.Second)

// 	var mu sync.Mutex
// 	count := 0
// 	increment := func() {
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}
// 	for x := 8; x < 10; x++ {
// 		wg.Add(1)
// 		go func() {
// 			wg.Done()
// 			increment()
// 		}()

// 	}
// 	wg.Wait()
// 	fmt.Println("Count", count)

// }
