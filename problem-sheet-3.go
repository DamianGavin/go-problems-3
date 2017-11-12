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

	fmt.Println("\nI am happy.")
	fmt.Println(elizaResponse("I am happy."))

	fmt.Println("\nI am not happy with your responses.")
	fmt.Println(elizaResponse("I am not happy with your responses."))

	fmt.Println("\nI am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(elizaResponse("I am not sure that you understand the effect that your questions are having on me."))

	fmt.Println("\nI am supposed to just take what you're saying at face value?")
	fmt.Println(elizaResponse("I am supposed to just take what you're saying at face value?"))
}

var responseStrings = []string{ //hardcoded given responses
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

func elizaResponse(userInput string) string { //input string of type string

	rex := regexp.MustCompile("(?i)\\bfather\\b") //i=case insensitive,(?i):match remainder of pattern with i

	if rex.MatchString(userInput) { //if exists
		return "Why don't you tell me some more about your Father?"
	}

	rex = regexp.MustCompile("(?i)i am (.*)") //i=case insensitive,(?i):match remainder of pattern with i

	found := rex.FindStringSubmatch(userInput)

	if len(found) > 1 {
		responseStrings := "How do you know that you are %s?"

		return fmt.Sprintf(responseStrings, found[1])
	}
	return responseStrings[rand.Intn(len(responseStrings))]
}
