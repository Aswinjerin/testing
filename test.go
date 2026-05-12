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

// orders := []user_info{
// 	{Order_Id: "ORD1001", Customer: "Divine", Restaurant: "Burger Hub", Amount: 450, Status: "PLACED"},
// 	{Order_Id: "ORD1002", Customer: "Alex", Restaurant: "Pizza Town", Amount: 0, Status: "PLACED"},
// 	{Order_Id: "ORD1003", Customer: "", Restaurant: "Food Spot", Amount: 300, Status: "PLACED"},
// 	{Order_Id: "ORD1004", Customer: "Sarah", Restaurant: "Sushi Zen", Amount: 1200, Status: "PLACED"},
// 	{Order_Id: "ORD1005", Customer: "", Restaurant: "Burger Hub", Amount: 450, Status: "PLACED"},
// }
