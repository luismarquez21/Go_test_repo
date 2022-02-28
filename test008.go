//defer segundo ejemplo 

package main

import "fmt"

func endTime(timestamp string) {
	fmt.Println("Program ended at", timestamp)
}

func main(){
	time := "1 PM"

	defer endTime(time)

	time = "2 PM"
    fmt.Println("doing something")
	fmt.Println("main finished")
	fmt.Println("time is", time)

}