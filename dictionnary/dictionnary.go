package dictionnary

import (
	"fmt"
)

func Get(dictionnaire map[string]string, key string) {
	fmt.Println("Definition for", key, ":", dictionnaire[key])
}

func Add(dictionnaire map[string]string, key string, value string) {
	dictionnaire[key] = value
}

func Remove(dictionnaire map[string]string, key string) {
	delete(dictionnaire, key)
}

func List(dictionnaire map[string]string) {
	for word := range dictionnaire {
		fmt.Println("key : " + word, "value : " + dictionnaire[word])
	} 
}