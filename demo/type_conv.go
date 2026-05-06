package demo

import (
	"fmt"
	"strconv"
)

func ma() {
	// Integer to float
	var a int = 10
	var b float64 = float64(a)
	fmt.Println("Integer:", a)
	fmt.Println("Converted to float:", b)

	// Float to integer
	var x float64 = 9.78
	var y int = int(x)
	fmt.Println("Float:", x)
	fmt.Println("Converted to int:", y)

	// Integer to string
	var num int = 123
	var str string = fmt.Sprintf("%d", num)
	fmt.Println("Integer:", num)
	fmt.Println("Converted to string:", str)

	// String to integer
	var s string = "4.45"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("enter the valid input value", err)
		return
	}
	fmt.Println("String:", s)
	fmt.Println("Converted to int:", n)
}
