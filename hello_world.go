package main

import "fmt"

func list() {
	fmt.Println("\nYo this is the list function")
}

func main() {

	// assign string "yo" to a and print it
	a := "yo"
	fmt.Println(a)

	// assign types and print
	const b int32 = 12
	fmt.Println("b:", b)
	fmt.Printf("b-type: %T ", b)

	list()

}
