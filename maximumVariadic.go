package main

import "fmt"

func maximum(args...int) int {
	max := args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
	
}

func main(){

	fmt.Println(maximum(1,2,4,7,10,100,99,98,50,11,24,23,35,85,75,4,56))
}