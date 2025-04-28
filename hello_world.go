package main

import "fmt"

func list() {
	// Creates a list of ints and prints out the index and the value
	fmt.Println("\nYo this is the list function")
	var myList = [...]int32{1, 2, 3}
	// for i := 0; i < len(myList); i++ {
	// 	fmt.Println(i, myList[i])
	// }
	for i, nums := range myList {
		fmt.Println(i, nums)
	}
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
