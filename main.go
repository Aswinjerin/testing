package main

import (
	"fmt"
	"log"
	"os"
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
	file, err := os.Create("invalid_orders.log")
	if err != nil {
		fmt.Println("File not created")
	} else {
		log.SetOutput(file)
		log.Println("File created")
	}
	fmt.Println("Food Delivery application")
	var n int
	fmt.Print("Enter number of orders: ")
	fmt.Scan(&n)
	orders := make([]user_info, n)
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Order %d\n", i+1)
		fmt.Print("Order ID: ")
		fmt.Scan(&orders[i].Order_Id)

		fmt.Print("Customer Name: ")
		fmt.Scan(&orders[i].Customer)

		fmt.Print("Restaurant Name: ")
		fmt.Scan(&orders[i].Restaurant)

		fmt.Print("Amount: ")
		fmt.Scan(&orders[i].Amount)

		fmt.Print("Status: placed / hold: ")
		fmt.Scan(&orders[i].Status)
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
				if i.Order_Id == "" || i.Customer == "" || i.Amount <= 0 || i.Status == "hold" {
					wrongorder++
					log.SetOutput(file)
					log.Printf("worker %d processed %+v is invalid \n", val, i.Order_Id)
				} else {
					conformorder++
				}
				totamt += i.Amount
			}
		}(j)
		//invalid_orders = append(wrongorder, invalid_orders.log)
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
