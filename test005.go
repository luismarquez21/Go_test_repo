// Ejemplo del uso de una función en Go
package main

import "fmt"

const Pi = 3.1416 

func area (radio float64) float64 {
	return Pi * radio * radio 
}

func dosomething(){
	fmt.Println("Hello World!")
}
//Imprimir mensaje de saludo
func greet(user string){
	fmt.Println("Hello " + user)
}
//Suma dos enteros
func add(a int, b int){
	c := a + b
	fmt.Println(c)
}

// Valor de retorno
func add2(a, b int) int64 {
	return int64(a + b)
}

// Múltiples valores de retorno
func addMult(a, b int) (int, int) {
	return a + b, a * b
}

func addMult2(a, b int)(add int, mul int) {
	add = a + b 
	mul = a * b 

	return // necessary
}









func main() {
	fmt.Println("El area de un circulo cuyo radio es 3 es: ", area(3))
	dosomething()
	greet("John Doe")
	add(1, 5)
	result := add2(1, 5)
	fmt.Println(result)
	addRes, multRes := addMult(2, 5)
	fmt.Println(addRes, multRes)
	_, mulRes := addMult(2, 5)
	fmt.Println(mulRes)
	addRes_two, multRes_two := addMult2(2, 5)
	fmt.Println(addRes_two, multRes_two)


}
