//Manejando errores en Scanf y Scan

package main

import "fmt"



func main(){
	var (
		nombre     string
		apellido   string
		telefono   int
		argumentos int
		err        error
	)
	argumentos, err = fmt.Scanf("%s %s %d", &nombre, &apellido, &telefono)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Printf("Todo bien, recibimos %d argumentos: %s, %s, %d", argumentos, nombre, apellido, telefono)
	}

}