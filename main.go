package main

import (
	"fmt"
	"github.com/hannanmiah/golang-tutorial/functions"
)


func main() {
	fmt.Println(functions.Sum(1,2))
	fmt.Println(functions.Sqrt(2))
	fmt.Println(functions.Sqrt(-3))
	fmt.Println(functions.VariadicSum(1,2,3,4,5,6,7,8,9,10))

	nums := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(functions.VariadicSum(nums...))
}

func init() {
    
}