
package main

import (
	"fmt"	// fmt.Println
	"os"	// os.Args
	"strconv"	// strconv.Atoi
)

func main() {
	x := 10

	if (x > 0) {
		fmt.Println("x is positive")
	} else if (x < 0) {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}
}
