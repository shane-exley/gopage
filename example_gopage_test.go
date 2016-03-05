package gopage_test

import (
	"fmt"
	"github.com/goibibo/gopage"
)

func ExamplePaginator() {
	p, _ := NewPaginator([]int{1, 2, 3, 4, 5})
	p.SetPageSize(2)

	fmt.Println(p.Page(1))
	fmt.Println(p.Page(2))
	fmt.Println(p.Page(3))
}
