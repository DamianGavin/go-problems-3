//Problem sheet 3 Data rep & query
//Damian Gavin 11/11/17

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano()) //generate a different responseStrings

	fmt.Println("People say I look like both my mother and my father")
	fmt.Println(ElizaResponse("People say I look like both my mother and my father"))

	fmt.Println("\nFather was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))

	fmt.Println("\nI was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))

	fmt.Println("\nI'm looking forward to the weekend.")
	fmt.Println(ElizaResponse("I'm looking forward to the weekend."))

	fmt.Println("\nMy grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
}

var responseStrings = []string{ //hardcode given responses
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

func ElizaResponse(input string) string {
	return responseStrings[rand.Intn(len(responseStrings))]
}
