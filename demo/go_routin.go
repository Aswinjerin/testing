package demo

import (
	"fmt"
)

var num int = 10

func increment(a int) int {
	a = num + 1
	fmt.Println(a)
	return a
}
func init() {
	increment(num)
}
func rout() {

	var arr = []int{1, 2, 3, 4}
	slc1 := []int{4, 5, 6, 78, 9}
	slc1 = append(slc1, 23, 5)
	slc1[3] = 43
	//slc1 = slc1[2:5]
	i := "hello"
	fmt.Println("number", num)
	fmt.Println("slice", slc1)

	fmt.Println("array", arr)
	fmt.Println(i)

}
