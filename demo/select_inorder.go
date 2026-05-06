package demo

import (
	"fmt"
	"time"
)

type Msg struct {
	index int
	text  string
}

func man() {
	ch1 := make(chan Msg)
	ch2 := make(chan Msg)
	ch3 := make(chan Msg)
	ch4 := make(chan Msg)
	ch5 := make(chan Msg)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- Msg{1, "Message from channel 1"}
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- Msg{2, "Message from channel 2"}
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch3 <- Msg{3, "Message from channel 3"}
	}()
	go func() {
		time.Sleep(4 * time.Second)
		ch4 <- Msg{4, "Message from channel 4"}
	}()
	go func() {
		time.Sleep(8 * time.Second)
		ch5 <- Msg{5, "Message from channel 5"}
	}()

	results := make(map[int]string)
	expected := 1

	for i := 0; i < 5; i++ {
		select {
		case m := <-ch1:
			results[m.index] = m.text
		case m := <-ch2:
			results[m.index] = m.text
		case m := <-ch3:
			results[m.index] = m.text
		case m := <-ch4:
			results[m.index] = m.text
		case m := <-ch5:
			results[m.index] = m.text
		}

		// print in correct order
		for {
			if val, ok := results[expected]; ok {
				fmt.Println(val)
				delete(results, expected)
				expected++
			} else {
				break
			}
		}
	}

}
