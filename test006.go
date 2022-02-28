//Función recursiva 
package main 

import "fmt"

// n! = nx(n-1)! where n > 0 
func getFactorial(num int) int {
	if num > 1 {
		return num * getFactorial(num-1) // 5! = 5x(n-1) where n > 0 : 5! = 1*2*3*4*5 = 5*4 = 20, 20*3 = 60, 60*2 = 120, 120*1 = 120

	}

	return 1 // 1! == 1
}

func main() {
	f := getFactorial(5)
	fmt.Println(f)
}