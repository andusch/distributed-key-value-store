package store

import (
	"fmt"
	"sync"
)

// pluggable compression
type Codec interface {
	Encode(string) string
	Decode(string) string
}

// for no compression
type NoOpCodec struct{}

func (n NoOpCodec) Encode(s string) string { return s }
func (n NoOpCodec) Decode(s string) string { return s }

type Store struct {
	mu    sync.RWMutex // Protects the map from concurrent access
	data  map[string]string
	codec Codec
}

func NewStore(codec Codec) *Store {
	if codec == nil {
		codec = NoOpCodec{}
	}
	return &Store{
		data:  make(map[string]string),
		codec: codec,
	}
}

func (s *Store) Get(key string) (string, bool) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	val, exists := s.data[key]

	if !exists {
		return "", false
	}
	return s.codec.Decode(val), true

}

func (s *Store) Put(key string, value string) (replaced bool) {

	s.mu.Lock()
	defer s.mu.Unlock()

	encoded := s.codec.Encode(value)
	_, replaced = s.data[key]
	s.data[key] = encoded

	action := "added"
	if replaced {
		action = "replaced"
	}

	fmt.Printf("Key %s %s (compressed: %d -> %d bytes)\n", key, action, len(value), len(encoded))

	return replaced

}

func (s *Store) Delete(key string) bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[key]; exists {
		delete(s.data, key)
		fmt.Println("Key", key, "was deleted.")
		return true
	}
	fmt.Println("Key", key, "does not exist.")
	return false

}

func (s *Store) Iterate() {

	s.mu.RLock()
	defer s.mu.RUnlock()

	fmt.Println("--- Store Contents ---")
	for key, val := range s.data {
		decoded := s.codec.Decode(val)
		fmt.Printf("Key: %s | Raw: %s | Stored: %s\n", key, decoded, val)
	}
	fmt.Println("---------------------")

}
