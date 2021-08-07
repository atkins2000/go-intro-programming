package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string // functionName(arguments by commas)(retrun type by commas)
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting (eb englishBot){
//	fmt.Println(eb.getGreeting())
//}

//func printGreeting (sb spanishBot){
//	fmt.Println(sb.getGreeting())
//}

func (englishBot) getGreeting() string { //reciever func but no use of eb/sb variable, hence no need to specify
	return "Hi there"
}

func (spanishBot) getGreeting() string { //reciever func but no use of eb/sb variable, hence no need to specify
	return "Hola"
}
