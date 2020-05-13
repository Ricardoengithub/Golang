package main 

import "fmt"

func fibonnaci(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibonnaci(x - 1)  + fibonnaci(x -2)
}

func main()  {
	fmt.Print("Enter a number: ")
	var x int
	fmt.Scanf("%d", &x)

	fmt.Println(fibonnaci(x))
}