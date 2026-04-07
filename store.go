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

func (store *Store) Put(key string, value string) {

	if store.M == nil {
		store.M = make(map[string]string)
	}

	store.M[key] = value

}

func (store Store) Get(key string) string {
	return store.M[key]
}
