package main

import (
	"fmt"
	"Dictionnaire/dictionnary"
)

func main() {
	dictio := dictionnary.NewDictionary("dictionary.json")

	dictio.Add("TESTEcrasement", "donnée")
	dictio.Add("NewFile", "ecrasementFile")

	fmt.Println("Dictionary before operations:")
	dictio.List()

	// Add a new entry
	dictio.Add("key3", "Definition 3")

	// Get the definition for key3
	fmt.Println("Definition for key3:", dictio.Get("key3"))

	// Remove key3
	dictio.Remove("key3")

	fmt.Println("Dictionary after operations:")
	dictio.List()
}