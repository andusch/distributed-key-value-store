package main

import "fmt"

func main() {

	store := NewStore()

	store.Put("1", "Andu")
	store.Put("2", "Andrei")
	store.Put("3", "Mama")
	store.Put("4", "Tata")

	store.Iterate()

	fmt.Println("==============================")

	store.Put("1", "Andu Scheusan")

	store.Iterate()

}
