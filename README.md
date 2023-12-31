# Repository GoLang Back-End Course

[![My Skills](https://skillicons.dev/icons?i=go&perline=1)](https://skillicons.dev)

Ce repository a été créé par Mateo Oudart pour le cours sur le langage back-end Go dispensé par Mr. Aziz Mounir. Ce cours fait partie de la 4ème année d'étude en développement web et mobile, dans le parcours de l'école Estiam : WMD (Web & Mobile Development).

## Postman Workspace to test the API : 

https://red-satellite-839868.postman.co/workspace/Dictionary-API~f068e27a-1efa-48e7-b8e1-50682fc52f8a/request/30673310-7b4d4bba-d4e3-4b1a-a0c9-67702fe7648c

## Comment Installer le Repository

1. **Clonez le Repository :**

   ```bash
   git clone https://github.com/username/nom-du-repository.git
   cd nom-du-repository

   ```

1. **Lancez l'application :**
   ```bash
   go run main.go
   ```

## Routes du Repository

- 🔄 / : Route principale, retourne "Hello, World!"
- 📖 /get : Route pour obtenir une définition.
- ➕ /add : Route pour ajouter une nouvelle entrée.
- 🗑️ /remove : Route pour supprimer une entrée.
- 📋 /list : Route pour lister toutes les entrées du dictionnaire.

## Contenu du Repository

- /dictionnary: Contient le code source du package dictionnaire.
- dictionnary.go: Définition et méthodes pour le dictionnaire.
- main.go: Fichier principal de l'application.
- dictionary.json: Fichier de stockage des données du dictionnaire.
