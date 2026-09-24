package main

import (
	"fmt"

	"github.com/kennedy1030/golangtest/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
