package main

import "fmt"

func main(){

	fmt.Print("Enter your name: ")
	var input string
	fmt.Scanf("%s", &input)

	fmt.Println("Hello " + input + "!")
}
