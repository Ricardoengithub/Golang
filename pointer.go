package main 

import "fmt"

func zero(x int) {
	x = 0
}

func zeroo(xPrt *int) {
	*xPrt = 0
}


func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	zeroo(&x)
	fmt.Println(x)
}