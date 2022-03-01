//Condicional if/else en Go

package main

import "fmt"

func main() {

	isLearningGo := true

	if isLearningGo {
		fmt.Println("¡Genial!, Go es un lenguaje increible 😎")
	} else {
		fmt.Println("¿Que esperas?, ¡Empieza a aprenderlo! 👨🏽‍💻")
	}

   // only if

   var edad int = 18

   if edad >= 18 {
	   fmt.Println("Eres mayor de edad")
   }

   var edad2 int = 16

   if edad2 >= 18 {
    fmt.Println("Eres mayor de edad")
   } else {
	   fmt.Println("Eres menor de edad")
   }

   if edad2 >=16 && edad2 <= 67 {
	   fmt.Println("Estas en edad de trabajar")
   }

   edad3 := 11

   if edad3 >= 16 && edad3 <= 67 {
	   fmt.Println("Estas en edad de trabajar")
   } else if edad3 < 16 {
	   fmt.Println("Eres aun muy joven para trabajar")
   } else {
	   fmt.Println("Estas ya Jubilado")
   }





}
