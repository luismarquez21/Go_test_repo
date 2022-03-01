//Variables
// Cómo asignar múltiples variables en una sola línea en Go
package main

import "fmt"

func main() {
	var nombre, apellido string = "Carmen", "Electra"
	fmt.Println(nombre, apellido)

	var (
		nombre2     string = "Paul"
		edad2       int    = 32
		pensionado2 bool   = false
	)
	fmt.Println("Nombre: ", nombre2)
	fmt.Println("Edad: ", edad2)
	fmt.Println("Pensionado: ", pensionado2)

}
