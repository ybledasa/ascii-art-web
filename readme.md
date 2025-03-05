# ASCII Art Web


## Description
ASCII-art-web est un serveur web en Go qui permet aux utilisateurs de générer de l'art ASCII en utilisant différents styles de bannières via une interface utilisateur graphique web.

## Auteurs

-ybledasa

## Utilisation
Installez Go : https://golang.org/doc/install
Clonez le dépôt : https://academy.digifemmes.com/git/gdimble/ascii-art-web.git
Accédez au répertoire du projet : cd ascii-art-web
Lancez le serveur : go run main.go
Ouvrez votre navigateur et allez à http://localhost:3001

## Détails d'Implémentation
Le serveur est implémenté en Go en utilisant le package standard net/http. Les modèles HTML sont stockés dans le répertoire templates.
## ERREUR 500
http.StatusInternalServerError

## ERREUR 404
http.NotFound(w, r)
## http://localhost:3001/5 

## ERREUR 400
w.WriteHeader(http.StatusBadRequest)
## le port: http://localhost:3001 
## http://localhost:3001/5 