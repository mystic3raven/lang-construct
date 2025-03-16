package main

import (
	"fmt" // fmt.Println
)

func divide(a, b int) (int, err) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

// Output:
