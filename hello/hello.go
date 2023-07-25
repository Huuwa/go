package main

import (
	"fmt"
	
	"github.com/huuwa/go/greetings"
)

func main(){
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}