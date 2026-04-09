package main

import "fmt"

func main() {

	store := NewStore()

	store.Put("1", "John")
	store.Put("2", "David")
	store.Put("3", "Mary")
	store.Put("4", "Jesus")

	store.Iterate()

	fmt.Println("==============================")

	store.Put("1", "God")

	store.Iterate()

	fmt.Println("==============================")

	store.Delete("2")

	store.Iterate()

	fmt.Println("==============================")

	original := "AABBBCCCC"
	compressed := run_length_encode(original)
	decoded := run_length_decode(compressed)
	fmt.Println("Original:", original)
	fmt.Println("Compressed:", compressed)
	fmt.Println("Decoded:", decoded)

}
