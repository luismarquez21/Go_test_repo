/* Una función en Go también es un tipo. Si dos funciones aceptan 
los mismos parámetros y  devuelven los mismos valores, 
entonces estas dos funciones son del mismo tipo.*/

package main 

import "fmt"

func add (a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b 
}

func calc(a int, b int, f func(int, int) int) int {
	r := f(a, b)
	return r 
}

func main() {
	fmt.Printf("Type of function add is       %T\n", add)
	fmt.Printf("Type of function subtract is  %T\n", subtract)
    addResult := calc(5, 3, add)
	subResult := calc(5, 3, subtract)
	fmt.Println("5+3 =", addResult)
	fmt.Println("5-3 =", subResult)
}