package main

import (
	"fmt"
	"sync"
	"time"
)

type user_info struct {
	Order_Id   string `json:"order_id"`
	Customer   string `json:"customer"`
	Restaurant string `json:"restaurant"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
}

func main() {
	fmt.Println("Food Delivery application")
	orders := []user_info{
		{Order_Id: "ORD1001", Customer: "Divine", Restaurant: "Burger Hub", Amount: 450, Status: "PLACED"},
		{Order_Id: "ORD1002", Customer: "Alex", Restaurant: "Pizza Town", Amount: 0, Status: "PLACED"},
		{Order_Id: "ORD1003", Customer: "", Restaurant: "Food Spot", Amount: 300, Status: "PLACED"},
		{Order_Id: "ORD1004", Customer: "Sarah", Restaurant: "Sushi Zen", Amount: 1200, Status: "PLACED"},
		{Order_Id: "ORD1005", Customer: "", Restaurant: "Burger Hub", Amount: 450, Status: "PLACED"},
	}
	ch := make(chan user_info, len(orders))
	var wg sync.WaitGroup
	var wrongorder int
	var conformorder int
	var totorder int
	var totamt int
	routine := len(orders)
	for j := range routine {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			for i := range ch {
				totorder++
				if i.Order_Id == "" || i.Customer == "" || i.Amount <= 0 {
					wrongorder++
					fmt.Printf("worker %d processed %v \n", val, i.Order_Id)
				} else {
					conformorder++
				}
				totamt += i.Amount
			}
		}(j)
	}
	for _, order := range orders {
		ch <- order
	}
	close(ch)
	wg.Wait()
	fmt.Println("")
	fmt.Println("------After Processing---------")
	fmt.Println("Invalid Orders   :", wrongorder)
	fmt.Println("Valid Orders     :", conformorder)
	fmt.Println("Total Orders     :", totorder)
	fmt.Println("Total Revenue    :", totamt)
	time.Sleep(2 * time.Millisecond)
}
