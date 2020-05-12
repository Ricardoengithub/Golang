package main

import "fmt"

func f() (a int, b int){
	return 10, 100
}

func main()  {
	x, y := f()

	w, error := f()

	z, ok := f()

	fmt.Println(x)
	fmt.Println(y)
}