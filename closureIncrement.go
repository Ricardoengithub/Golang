package main 

import "fmt"

func main(){
    	increment := func(x int) int{
		x++
		return x
	}

	fmt.Println(increment(2))
	fmt.Println(increment(5))
}