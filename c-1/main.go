package main;

import "fmt"

func main() {
	var username string = "waggslane"
	var password string = "100000"
	fmt.Println("Authorization: Basic",username + ":" + password)
}