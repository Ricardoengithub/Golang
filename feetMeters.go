package main

import "fmt"

func main(){

	fmt.Print("Enter feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	feet = feet * 0.3048

	fmt.Printf("Meters: %.2f \n", feet)
}
