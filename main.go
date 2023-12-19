package main

import (
	"fmt"
)

func main() {

	dictio := make(map[string]string)
	dictio["key1"] = "10"
	dictio["key2"] = "12"

	fmt.Println(dictio)

	add(dictio, "key3", "13")
	get(dictio, "key3")
	remove(dictio, "key3")

	fmt.Println(dictio)

	List(dictio)
}

func get(dictionnaire map[string]string, key string) {
	fmt.Println("Definition for", key, ":", dictionnaire[key])
}

func add(dictionnaire map[string]string, key string, value string) {
	dictionnaire[key] = value
}

func remove(dictionnaire map[string]string, key string) {
	delete(dictionnaire, key)
}

func List(dictionnaire map[string]string) {
	for word := range dictionnaire {
		fmt.Println("key : " + word, "value : " + dictionnaire[word])
	} 
}