package main

import "fmt"

func factorial(x int) int {
	
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}


func main(){
	fmt.Print("Escribe un número: ")
	var x int
	fmt.Scanf("%d", &x)
	fmt.Println(factorial(x))
}