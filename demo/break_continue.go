package demo

import "fmt"

func mn() {
	for i := 1; i <= 5; i++ {

		if i == 3 {
			break
		}

		fmt.Println("break loop", i)
	}
outer:
	for h := 0; h <= 5; h++ {
		for j := 1; j <= 5; j++ {

			if j == 3 {
				continue outer
			}
			fmt.Println("continue loop", j, h)
		}
	}

}
