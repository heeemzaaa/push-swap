package main

import (
	"fmt"
	p "push/push-swap"
)

func main() {
	s1 := []int{0, 11, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := []int{}
	v1 := p.Push_Swap(s1, s2)
	fmt.Print(v1)
}
