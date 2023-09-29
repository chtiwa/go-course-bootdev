package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {

	return fmt.Sprintln("Authorization: Basic " + authI.username + ":" + authI.password)
}

func testing(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	testing(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	testing(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	testing(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
