package dictionnary

import (
	"fmt"
)

/// Object Structure
type Dictionary struct {
	entries map[string]string
}

///Constructor
func NewDictionary() *Dictionary {
	return &Dictionary{
			entries: make(map[string]string),
	}
}

///Functions
func (d *Dictionary) Get(key string) string {
	return d.entries[key]
}

func (d *Dictionary) Add(key, value string) {
	d.entries[key] = value
}

func (d *Dictionary) Remove(key string) {
	delete(d.entries, key)
}

func (d *Dictionary) List() {
	for word := range d.entries {
		fmt.Println("key:", word, "value:", d.entries[word])
	}
}