package main

import (
	"fmt"
	"Dictionnaire/dictionnary"
)

func main() {

	dictio := make(map[string]string)
	dictio["key1"] = "10"
	dictio["key2"] = "12"

	fmt.Println(dictio)

	dictionnary.Add(dictio, "key3", "13")
	dictionnary.Get(dictio, "key3")
	dictionnary.Remove(dictio, "key3")

	fmt.Println(dictio)

	dictionnary.List(dictio)
}