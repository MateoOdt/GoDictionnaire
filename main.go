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
	fmt.Println(dictio)

}

func add(dictionnaire map[string]int, key string, value int) {
	dictionnaire[key] = value
}