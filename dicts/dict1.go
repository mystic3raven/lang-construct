package main // package name
import (
	"fmt" // fmt.Println
)

func main() {

	person := map[string]interface{}{"name": "Alice", "age": 30} // map of string to interface
	fmt.Println(person["name"])                                  // Alice
	fmt.Println(person["age"])                                   // 30
}
