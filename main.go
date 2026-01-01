package main

import (
	"fmt"
	"rsc.io/quote"
	"orion.com/go-basic/greeting"
)



func main(){
	fmt.Println("hello, let's start go.")
	fmt.Println(quote.Hello())
	hello()
	

	returnMsg := greeting.SayGreeting("Orion") 
	fmt.Println(returnMsg)
}