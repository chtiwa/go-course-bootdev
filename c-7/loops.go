package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	greetings("chtiwa")
	// 	fmt.Println()
	// 	break
	// }

	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

}

func greetings(name string) {
	fmt.Printf("greetings %v", name)
}
