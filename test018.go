//Convertir entre tipos
//Esto se logra utilizando la librería strconv que provee muchos métodos de conversión

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	var mayorDeEdad string = "true"
	boolVal, _ := strconv.ParseBool(mayorDeEdad)
	fmt.Println(boolVal, reflect.TypeOf(boolVal))
	var mayorDeEdad2 bool = true
    strVal := strconv.FormatBool(mayorDeEdad2)
    fmt.Println(strVal, reflect.TypeOf(strVal))


	cadena1 := "48"
    cadena2 := "52"
    var x1, x2 int
    x1, _ = strconv.Atoi(cadena1)
    x2, _ = strconv.Atoi(cadena2)
    suma := x1 + x2
    fmt.Println("La suma de los dos valores es:", suma)
}