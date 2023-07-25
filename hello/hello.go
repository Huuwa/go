package main

import (
	"fmt"
	
	"github.com/Huuwa/go/greetings"
)

func main(){
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}