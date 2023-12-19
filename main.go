package main

import (
	"fmt"
	"Dictionnaire/dictionnary"
)


func main() {

	dictio := dictionnary.NewDictionary()

	dictio.Add("key1", "10")
	dictio.Add("key2", "12")

	fmt.Println("Dictionary before operations:")
	dictio.List()

	// Add a new entry
	dictio.Add("key3", "13")

	// Get the value for key3
	fmt.Println("Definition for key3:", dictio.Get("key3"))

	// Remove key3
	dictio.Remove("key3")

	fmt.Println("Dictionary after operations:")
	dictio.List()
}