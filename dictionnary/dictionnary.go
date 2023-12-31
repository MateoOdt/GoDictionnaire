package dictionnary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Entry represents a dictionary entry with a name and definition
type Entry struct {
	Nom        string `json:"nom"`
	Definition string `json:"definition"`
}

// Dictionary structure
type Dictionary struct {
	entries []Entry
	file    string
}

// NewDictionary creates a new Dictionary instance with the specified filename
func NewDictionary(filename string) *Dictionary {
	d := &Dictionary{
		entries: make([]Entry, 0),
		file:    filename,
	}
	d.loadFromFile()
	return d
}

// Get returns the definition for the given name
func (d *Dictionary) Get(nom string) string {
	for _, entry := range d.entries {
		if entry.Nom == nom {
			return entry.Definition
		}
	}
	return ""
}

// Add adds a new entry to the dictionary
func (d *Dictionary) Add(nom, definition string) {
	go func() {
		entry := Entry{Nom: nom, Definition: definition}
		d.entries = append(d.entries, entry)
		d.saveToFile()
	}()
} 

// Remove removes an entry from the dictionary
func (d *Dictionary) Remove(nom string) {
	go func() {
		for i, entry := range d.entries {
			if entry.Nom == nom {
				d.entries = append(d.entries[:i], d.entries[i+1:]...)
				break
			}
		}
		d.saveToFile()
	}()
}

// List displays all entries in the dictionary
func (d *Dictionary) List() []Entry {
	return d.entries
}

///Endpoints :
func HandleWelcomeRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func HandleGet(d *Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nom := r.URL.Query().Get("nom")
		definition := d.Get(nom)
		response := map[string]string{"nom": nom, "definition": definition}
		json.NewEncoder(w).Encode(response)
	}
}

func HandleAdd(d *Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entry Entry
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		d.Add(entry.Nom, entry.Definition)
		w.WriteHeader(http.StatusCreated)

		///Return Successfull message on Add Entry
		response := "Entry Successfull"
		json.NewEncoder(w).Encode(response)
	}
}

func HandleRemove(d *Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nom := r.URL.Query().Get("nom")
		d.Remove(nom)
		w.WriteHeader(http.StatusOK)
	}
}

func HandleList(d *Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entries := d.List()
		json.NewEncoder(w).Encode(entries)
	}
}

func (d *Dictionary) loadFromFile() {
	fileContent, err := ioutil.ReadFile(d.file)
	if err == nil {
		json.Unmarshal(fileContent, &d.entries)
	}
}

func (d *Dictionary) saveToFile() {
	fileContent, err := json.MarshalIndent(d.entries, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	err = ioutil.WriteFile(d.file, fileContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
