package main

import "fmt"

func main(){
	
	fmt.Print("F: ")
	var f float64
	fmt.Scanf("%f", &f)

	f = (f - 32) * 5 / 9

	fmt.Printf("C: %.2f\n", f)
}
