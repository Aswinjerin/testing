package demo

import (
	"context"
	"fmt"
	"time"
)

func pass(ps context.Context) {
	for {
		select {
		case <-ps.Done():
			fmt.Println("Process stopped:", ps.Err())
			return
		default:
			fmt.Println("work on process ⚡")
			time.Sleep(1 * time.Second)
		}
	}
}
func in() {
	ps, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go pass(ps)

	time.Sleep(1 * time.Second)
	fmt.Println("Work Finish")
	if deadline, ok := ps.Deadline(); ok {
		fmt.Println("deadline occures at : ", deadline)
	}

}

// how to debug if we got a error in the golang program

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context) {

// 	for {

// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Goroutine stopped:", ctx.Err())
// 			return
// 		default:
// 			fmt.Println("Goroutine working...")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
// 	defer cancel()

// 	go worker(ctx)

// 	time.Sleep(8 * time.Second)

// 	fmt.Println("Main function done")

// }
