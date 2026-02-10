package main

import (
	"fmt"

	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Pujan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	messages, err := greetings.Hellos([]string{"John", "Jacob", "Tony"})
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}
