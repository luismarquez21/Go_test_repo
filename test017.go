// Verificar tipos de datos con Go
//Esto se puede hacer utilizando el paquete reflect 
 package main

 import (
	 "fmt"
	 "reflect"
 )

 func main () {
	 var alumno string = "Jose luis"
	 var edad int = 22
	 var peso float64 = 85.5 
	 fmt.Println(reflect.TypeOf(alumno))
	 fmt.Println(reflect.TypeOf(edad))
	 fmt.Println(reflect.TypeOf(peso))
 }
