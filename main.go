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

/// Gestionnaire de requêtes pour le point de terminaison POST /ascii-art
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.URL.Path == "/ascii-art" {
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if strings.TrimSpace(text) == "" || banner == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

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

		asciiArt := generateASCIIArt(text, banner)

		// Utilise index.html avec le message à afficher
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Erreur lors du chargement de la page.", http.StatusInternalServerError)
			return
		}

		form := Form{
			Text:    r.FormValue("text"), // pour garder le texte original avec \n
			Banner:  banner,
			Message: asciiArt,
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, form)
	} else {
		http.NotFound(w, r)
	}
}
func generateASCIIArt(text, banner string) string {
	// Lire le fichier de bannière correspondant
	asciiArtTemplate, err := ioutil.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Erreur de lecture du fichier de bannière :", err)
		return ""
	}

	// Gérer les retours à la ligne personnalisés \n
	userInput := strings.ReplaceAll(text, "\\n", "\n")

	// Vérification des caractères valides (ASCII imprimables ou saut de ligne)
	for _, ch := range userInput {
		if (ch < 32 || ch > 126) && ch != '\n' {
			fmt.Println("Caractère non autorisé détecté :", string(ch))
			return ""
		}
	}

	// Chaque caractère (de ' ' à '~') correspond à une section dans le fichier modèle
	asciiBlocks := strings.Split(string(asciiArtTemplate)[1:], "\n\n")
	if len(asciiBlocks) != 95 {
		fmt.Println("Fichier modèle corrompu ou mal formaté.")
		return ""
	}

	// Construction de l'art ASCII ligne par ligne
	var asciiArtBuilder strings.Builder
	lines := strings.Split(userInput, "\n")

	for _, line := range lines {
		if line == "" {
			asciiArtBuilder.WriteString("\n")
			continue
		}

		// Il y a généralement 8 lignes par caractère
		for row := 0; row < 8; row++ {
			for _, ch := range line {
				block := strings.Split(asciiBlocks[ch-32], "\n")
				if row < len(block) {
					asciiArtBuilder.WriteString(block[row])
				}
			}
			asciiArtBuilder.WriteString("\n")
		}
	}

	return asciiArtBuilder.String()
}
