package main

import (
	"fmt"
	"time"
)

func main() {
	var a int
	fmt.Print("Enter the number ")
	fmt.Scan(&a)
	val := a / 10
	go func(x int) {
		for i := 0; i <= val; i++ {
			x = 0 + i
		}
		fmt.Println("1st val ", x)
	}(val)
	val2 := val * 2
	go func(x int) {
		for i := val; i <= val2; i++ {
			x = i
		}
		fmt.Println("2nd val ", x)
	}(val2)
	val3 := val * 3
	go func(x int) {
		for i := val2; i <= val3; i++ {
			x = i
		}
		fmt.Println("3rd val ", x)
	}(val3)
	val4 := val * 4
	go func(x int) {
		for i := val3; i <= val4; i++ {
			x = 0 + i
		}
		fmt.Println("4th val ", x)
	}(val4)
	val5 := val * 5
	go func(x int) {
		for i := val4; i <= val5; i++ {
			x = 0 + i
		}
		fmt.Println("5th val ", x)
	}(val5)
	val6 := val * 6
	go func(x int) {
		for i := val5; i <= val6; i++ {
			x = 0 + i
		}
		fmt.Println("6th val ", x)
	}(val6)
	val7 := val * 7
	go func(x int) {
		for i := val6; i <= val7; i++ {
			x = 0 + i
		}
		fmt.Println("7th val ", x)
	}(val7)
	val8 := val * 8
	go func(x int) {
		for i := val7; i <= val8; i++ {
			x = 0 + i
		}
		fmt.Println("8th val ", x)
	}(val8)
	val9 := val * 9
	go func(x int) {
		for i := val8; i <= val9; i++ {
			x = 0 + i
		}
		fmt.Println("9th val ", x)
	}(val9)
	val10 := val * 10
	go func(x int) {
		for i := val9; i <= val10; i++ {
			x = 0 + i
		}
		fmt.Println("10th val ", x)
	}(val10)
	time.Sleep(2 * time.Second)
}
