package main// what should this be?

import (
	"fmt"
	"ccwhois/greetings"
)


// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}


func main() {
	fmt.Println(Hello("Varun"))
	fmt.Println(greetings.Hello("Varun"))
}