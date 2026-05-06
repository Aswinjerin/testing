package demo

import (
	"fmt"
	"strconv"
)

func main() {
	var n, sum int
	for i := 0; i < 5; i++ {
		fmt.Println("Enter the value : ")
		fmt.Scanf("%d\n", &n)
		sum = sum + n
	}
	fmt.Println(sum)
	var a int
	var str string
	var err error
	for {
		fmt.Println("Enter numb :")
		fmt.Scanf("%d\n", &a)
		a, err = strconv.Atoi(str)
		if err != nil {
			fmt.Println("BREAKING")
			break
		}
	}

	fmt.Println("AFTER BREAK PROCESSING")

	var arr [5]int
	arr = [...]int{2, 4, 5, 6, 8}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d:%d\n", i, arr[i])
	}
}
