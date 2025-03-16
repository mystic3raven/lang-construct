package main // package name

import (
	"fmt" // fmt.Println
	"os"
)

func main() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	err = os.WriteFile("file.txt", []byte("Hello, World!"), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
