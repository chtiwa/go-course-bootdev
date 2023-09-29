package main

import "fmt"

// variadic
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
	names := []string{"bob", "sue", "alice"}
	printStrings(names...)
}
