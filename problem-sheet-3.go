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

	fmt.Println("\nIm happy.")
	fmt.Println(elizaResponse("I am happy."))

	fmt.Println("\nI am not happy with your responses.")
	fmt.Println(elizaResponse("I am not happy with your responses."))

	fmt.Println("\nI'm not sure that you understand the effect that your questions are having on me.")
	fmt.Println(elizaResponse("I am not sure that you understand the effect that your questions are having on me."))

	fmt.Println("\ni am supposed to just take what you're saying at face value?")
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

	//this regex is to accomodate different ways of saying "I am"
	//1st capture: the ` is instead of ", (?i)matches remainder of pattern with i
	//modifier:case insensitive. The 2nd i is a literal character match and
	//again either case.(?:'|\sa)? is a non-capturing group. the ? match between 0-1
	//times, as many as possible, giving back as needed(greedily) The ' is any ', eg I'm
	//\s matches any whitespace. a is literal character, as is m. the 2nd capture (.*):
	//.*matches any character except full stops.
	//g modifier: all matches ie. not just 1st match

	rex = regexp.MustCompile(`(?i)i(?:'|\sa)?m (.*)`) //i=case insensitive,(?i):match remainder of pattern with i

	found := rex.FindStringSubmatch(userInput)

	if len(found) > 1 {
		responseStrings := "How do you know that you are %s?"

		return fmt.Sprintf(responseStrings, found[1])
	}
	return responseStrings[rand.Intn(len(responseStrings))]
}
