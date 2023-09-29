package main

import "fmt"

// variadic
func sum(nums ...int) int {
	// nums is a slice
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func main() {
	total := sum(1, 2, 3)
	fmt.Println(total)
}
