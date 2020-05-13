package main 

import "fmt"

func first() {
	fmt.Println("Primero")
}

func second() {
	fmt.Println("Segundo")
}

func main() {
	defer second()
	first()
}