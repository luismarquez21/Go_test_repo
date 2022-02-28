//Funciones en go
package main

import "fmt"

func RetornaUno(argumento, otroArgumento int) int {
	return 1
}

func RetornaUnoYDos(argumento2 , otroArgumento2 int) (int, int) {
	return 1,2
}

func main() {
	var_one := RetornaUno(1, 2)
	
	fmt.Println(var_one)
	
}
