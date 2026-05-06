package demo

import (
	"fmt"
	"time"
)

func mein() {
	// fmt.Println("Start")
	// panic("something went wrong!")
	// defer fmt.Println("End") // never runs

	//---------------------------------------

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	ch5 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- "Message from channel 3"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "Message from channel 4"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch5 <- "Message from channel 5"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- "Message from channel 2"
	}()
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)

		case msg2 := <-ch2:
			fmt.Println(msg2)
		case msg3 := <-ch3:
			fmt.Println(msg3)
		case msg4 := <-ch4:
			fmt.Println(msg4)
		case msg5 := <-ch5:
			fmt.Println(msg5)

		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: no message received in 2 seconds")
		default:
			fmt.Println("Nothing Ready")

		}
	}
	// ch3 := make(chan string)
	// ch4 := make(chan string)

	// fmt.Println(<-ch3)
	// fmt.Println(<-ch4)
}
