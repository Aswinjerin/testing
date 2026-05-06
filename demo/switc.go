package demo

import (
	"fmt"
)

func demo() {
	var n byte
	fmt.Println("Enter want to subscribe (y/n) :", n)
	fmt.Scanf("%c", &n)
	switch n {
	case 'y', 'Y':
		fmt.Println("Thankyou")
	case 'n', 'N':
		fmt.Println("sorry")
	default:
		fmt.Println("Choose one")
	}
}
