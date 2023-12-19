package main

import (
	"fmt"
)

func main() {

	dictio := make(map[string]int)
	dictio["key1"] = 10
	dictio["key2"] = 12

	fmt.Println(dictio)

	add(dictio, "key3", 13)
	get(dictio, "key3")
	remove(dictio, "key3")
	fmt.Println(dictio)

}

func get(dictionnaire map[string]int, key string) {
	fmt.Println(dictionnaire[key])
}

func add(dictionnaire map[string]int, key string, value int) {
	dictionnaire[key] = value
}

func remove(dictionnaire map[string]int, key string) {
	delete(dictionnaire, key)
}