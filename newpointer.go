package main 

import "fmt"

func one(xPrt *int){
	*xPrt = 1
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(xPtr)
}