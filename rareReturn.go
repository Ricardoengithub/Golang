package main

import "fmt"

func f2() (v int){
	r := 1
	v = r+1
	fmt.Println(r)
	return
}

func main(){
	fmt.Println(f2())
}