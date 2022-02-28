//defer palabra clave
package main 

import "fmt"

func sayDone() {
	fmt.Println("I am done")
}

func main() {
	fmt.Println("main started")

	defer sayDone()

	fmt.Println("main finished")

}
