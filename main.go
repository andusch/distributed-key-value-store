package main

import (
	"distributed-key-value-store/pkg/codec"
	"distributed-key-value-store/pkg/store"
	"fmt"
	"sync"
)

func main() {

	str := store.NewStore(codec.RLECodec{})

	str.Put("1", "John")
	str.Put("2", "David")
	str.Put("3", "Mary")
	str.Put("4", "Jesus")
	str.Iterate()

	fmt.Println("==============================")
	str.Put("1", "God")
	str.Iterate()

	fmt.Println("==============================")
	str.Delete("2")
	str.Iterate()

	fmt.Println("==============================")
	fmt.Println("Testing concurrency...")
	testConcurrency(str)

	c := codec.RLECodec{}
	encoded := c.Encode("AABBBCCCC")
	decoded := c.Decode(encoded)
	fmt.Printf("Original: AABBBCCCC\nEncoded: %s\nDecoded: %s\n", encoded, decoded)

}

func testConcurrency(s *store.Store) {

	var wg sync.WaitGroup
	n := 100

	// Concurrent writes
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id%10) // Some collisions
			val := fmt.Sprintf("value-%d", id)
			s.Put(key, val)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Completed %d concurrent writes\n", n)
}
