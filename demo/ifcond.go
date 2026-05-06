package demo

import (
	"fmt"
	"strconv"
)

func ifcond() {
	var a int
	fmt.Printf("Enter a value : ")
	fmt.Scanf("%d", &a)
	/*if n := a; a%2 == 0 {
		fmt.Printf("%d is a even number %d times", a, n)
	} else {
		fmt.Printf("%d is a odd number %d times", a, n)
	}*/
	fmt.Println(fizzBuss(a))
}
func fizzBuss(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(n)
	}
}
