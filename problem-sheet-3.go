//Problem sheet 3 Data rep & query
//Damian Gavin 11/11/17

package main

import (
	"math/rand"
)

func main() {

}

var responseStrings = []string{ //hardcode given responses
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

func ElizaResponse(input string) string {
	return responseStrings[rand.Intn(len(responseStrings))]
}
