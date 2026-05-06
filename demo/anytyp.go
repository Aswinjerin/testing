package demo

import (
	"fmt"
)

// Interval structure
func anytyp() {
	m := map[any]any{
		"id": 1,
		1:    "Book",
	}
	fmt.Println(m)
	type Item struct {
		ID    any
		Name  string
		Price float64
	}
	var name any
	name = "Hello interface"
	data := Item{ID: "aswin", Name: "Hello", Price: 220.45}
	fmt.Println(data)
	fmt.Println(name)
}
