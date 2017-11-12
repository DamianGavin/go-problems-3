//Problem sheet 3 Data rep & query
//Damian Gavin 11/11/17
//regex helped by https://stackoverflow.com/questions/9348326/regex-find-word-in-the-string
//and https://regex101.com/

package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano()) //generate a different responseStrings

	fmt.Println("People say I look like both my mother and my father")
	fmt.Println(elizaResponse("People say I look like both my mother and my father")) //depending on "elizaResponse" will change

	fmt.Println("\nFather was a teacher.")
	fmt.Println(elizaResponse("Father was a teacher."))

	fmt.Println("\nI was my father’s favourite.")
	fmt.Println(elizaResponse("I was my father’s favourite."))

	fmt.Println("\nI'm looking forward to the weekend.")
	fmt.Println(elizaResponse("I'm looking forward to the weekend."))

	fmt.Println("\nMy grandfather was French!")
	fmt.Println(elizaResponse("My grandfather was French!"))
}

var responseStrings = []string{ //hardcoded given responses
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

func elizaResponse(input string) string { //input string of type string

	rex := regexp.MustCompile("(?i)\\bfather\\b") //i=case insensitive

	if len(rex.FindStringIndex(input)) > 0 {
		return "Why don't you tell me some more about your Father?"
	}

	return responseStrings[rand.Intn(len(responseStrings))]
}
