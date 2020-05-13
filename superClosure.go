package main 

import "fmt"


func makeEvenGenerator() func() uint {
	i := uint(0)

	return func () (ret uint) {
		ret = i	
		i += 2 
		return 
	}
}


func makeOddGenerator() func() int {
	j := 1

	return func() (r int){
		r = j
		j += 2
		return
	}
}

func main(){
	
	nextEven := makeEvenGenerator()
	nextOdd := makeOddGenerator()


	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())


	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}