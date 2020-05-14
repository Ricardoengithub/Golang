package main

import "fmt"

func swap(x *int, y *int){
	*x, *y = *y, *x
}

func main() {
	x := 2
	y := 5 
	swap(&x,&y)
	fmt.Println(x)
	fmt.Println(y)
}