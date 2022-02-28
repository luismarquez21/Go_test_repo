// Arreglos (Arrays)
package main

import "fmt"

func main()  {
  var a [3]int // arreglo de 3 enteros
  fmt.Println(a[0]) // imprime el primer elemento
  fmt.Println(a[len(a)-1]) // imprime el último elemento, a[2]
  var ages [3]int = [3]int{10, 20, 30}
  var languages [3]string = [3]string{"go", "php"}
  fmt.Println(ages) // [10, 20, 30]
  fmt.Println(languages[0]) // ''
}