package main

import (
	"fmt"
	"html/template"
	"io/ioutil" 
	"log"
	"net/http"
	"strings"
)

// Structure pour stocker les données du formulaire
type Form struct {
	Text    string
	Banner  string
	Message string
}

func main() {
	// Gestionnaire de requêtes pour la page d'accueil
	http.HandleFunc("/", homeHandler)

	// Gestionnaire de requêtes pour le point de terminaison POST /ascii-art
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Serveur en écoute sur le port: http://localhost:3001.")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

// Gestionnaire de requêtes pour la page d'accueil (GET /)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/" {
		// Charger le fichier HTML de la page d'accueil
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Erreur lors du chargement de la page d'accueil.", http.StatusInternalServerError)
			return
		}

		// Afficher la page d'accueil
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, nil)

	} else {
		//generer erreur 404
		http.NotFound(w, r)
	}
}

// Gestionnaire de requêtes pour le point de terminaison POST /ascii-art
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.URL.Path == "/ascii-art" {
		// Extraire les données du formulaire
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == " " || banner == "" {
			//generer erreur 400
			w.WriteHeader(http.StatusBadRequest)
		} else {

			textRune := []rune{}

			for _, ltr := range text {
				if ltr == 10 {
					textRune = append(textRune, '\\')
					textRune = append(textRune, 'n')
				} else if ltr != 13 {
					textRune = append(textRune, ltr)
				}
			}
			text = string(textRune)
			fmt.Println(text)
			// Logique pour générer l'ASCII Art à partir des données du formulaire
			asciiArt := generateASCIIArt(text, banner)

			// Charger le fichier HTML du résultat
			tmpl, err := template.ParseFiles("templates/result.html")
			if err != nil {
				http.Error(w, "Erreur lors du chargement du résultat.", http.StatusInternalServerError)
				return
			}

			form := Form{
				Text:    text,
				Banner:  banner,
				Message: asciiArt,
			}

			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, form)
		}
	} else {
		http.NotFound(w, r)
	}
}

// Fonction pour générer l'ASCII Art
func generateASCIIArt(text, banner string) string {
	// Logique de génération de l'ASCII Art en fonction du texte et de la bannière sélectionnée
	// Remplacez cette fonction par votre propre logique de génération de l'ASCII Art

	// Exemple de génération d'ASCII Art simple

	// asciiArt := ""
	// for _, char := range text {
	// 	asciiArt += string(char) + "\n"
	// }

	// return asciiArt

	// Lecture et stockage du contenu du fichier modèle ASCII
	asciiArtTemplate, err := ioutil.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Remplacement des séquences "\n" par des sauts de ligne réels
	userInput := strings.ReplaceAll(text, "\\n", "\n")
	if userInput == "\n" {
		fmt.Println()
		return ""
	}

	// Vérification des caractères imprimables dans l'entrée utilisateur
	for i := 0; i < len(userInput); i++ {
		if (userInput[i] < 32 || userInput[i] > 127) && userInput[i] != 10 {
			fmt.Println("Incorrect input.")
			return ""
		}
	}

	// Division du modèle ASCII en blocs, un pour chaque caractère
	asciiBlocks := strings.Split(string(asciiArtTemplate)[1:], "\n\n")
	if len(asciiBlocks) != 95 {
		fmt.Println("Incorrect template.")
		return ""
	}

	// Traitement de l'entrée et création de l'art ASCII
	inputLines := strings.Split(userInput, "\n")
	asciiArtResult := ""
	for _, line := range inputLines {
		if line == "" && asciiArtResult != "" {
			asciiArtResult += string('\n')
			continue
		}

		// Construction de l'art ASCII pour chaque ligne
		for rowIndex := 0; rowIndex < 8; rowIndex++ {
			for i := 0; i < len(line); i++ {
				charBlock := strings.Split(asciiBlocks[line[i]-32], "\n")[rowIndex]
				asciiArtResult += charBlock
			}
			asciiArtResult += string('\n')
		}
	}
	return asciiArtResult
}
