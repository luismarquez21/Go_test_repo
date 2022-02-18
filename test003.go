//Ejemplo del uso del tipo booleano en Go
package main

import "fmt"

func main(){
	//boolean operators 
	a := 3 
	b := 4
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true 
	fmt.Println(a < b)  // true
    fmt.Println(a > b) // false 
	fmt.Println( a >= b ) // false
    fmt.Println( a <= b) // true 

	var doesUniverseExist bool
	doesUniverseExist = true
	fmt.Printf("Data = %v, Type = %T", doesUniverseExist, doesUniverseExist);

}