package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	fmt.Println(x[1:4])
	fmt.Println(x[4:])
	fmt.Println(x[:4])
	fmt.Println(x[1:4:6])
	fmt.Println(x[1:4:4])
	x.append(10)
	fmt.Println(x)
}
