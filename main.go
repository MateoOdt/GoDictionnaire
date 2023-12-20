package main

import (
	"Dictionnaire/dictionnary"
	"net/http"
)

func main() {
	dict := dictionnary.NewDictionary("dictionary.json")

	http.HandleFunc("/", dictionnary.HandleWelcomeRoot)

	http.HandleFunc("/get", dictionnary.HandleGet(dict))
	http.HandleFunc("/add", dictionnary.HandleAdd(dict))
	http.HandleFunc("/remove", dictionnary.HandleRemove(dict))
	http.HandleFunc("/list", dictionnary.HandleList(dict))

	http.ListenAndServe(":8080", nil)
}