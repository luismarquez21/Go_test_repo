// Entrada de texto Scan y Scanf
package main

import "fmt"



func main(){
	// scan
	/*var mensaje string
    fmt.Scan(&mensaje)
    fmt.Println(mensaje)
	// Scanf */
	var (
		nombre   string
		apellido string
		telefono int
	)
        // nombre apellido telefono
	fmt.Scanf("%s %s %d", &nombre, &apellido, &telefono)
	fmt.Printf("Nombre: %s, Apellido: %s, telefono: %d", nombre, apellido, telefono)

}