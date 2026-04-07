package main

import "fmt"

type Store struct {
	M map[string]string
}

func NewStore() *Store {
	return &Store{make(map[string]string)}
}

func (store *Store) Iterate() {

	for key, value := range store.M {
		fmt.Println("Key:", key, " Value:", value)
	}

}

func (store Store) Get(key string) string {
	return store.M[key]
}

func (store *Store) Put(key string, value string) {

	if store.M == nil {
		store.M = make(map[string]string)
	}

	if store.M[key] != "" {
		fmt.Println("Value", store.M[key], "was replaced with", value)
	} else {
		fmt.Println("Value", value, "was added in the store with key", key)
	}

	store.M[key] = value

}

func (store *Store) Delete(key string) {

	if _, exists := store.M[key]; exists {
		delete(store.M, key)
		fmt.Println("Key", key, "was deleted.")
	} else {
		fmt.Println("Key", key, "does not exist in the store.")
	}

}
