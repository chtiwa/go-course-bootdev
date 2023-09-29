package main

import "fmt"

// import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.message == "" {
		return false
	}
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func test(mToSend messageToSend) {
	if canSendMessage(mToSend) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("can't send message!")
	}
}

func main() {
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 0,
		},
	})
}
