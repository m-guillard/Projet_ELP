package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func gestion_erreur(err error, message string) {
	// Arrête le programme si il y a une erreur et affiche un message d'erreur dans le terminal
	if err != nil {
		fmt.Printf("Erreur : %s - %v\n", message, err)
		os.Exit(1) // Arrête le programme avec un code de retour 1 (indiquant une erreur)
	}
}

func nom_fichier(fichierId string, adrClient string) string {
	// Définit un fichier en fonction de la date et de l'adresse IP du client
	// Date du jour
	now := time.Now()
	// Formate la date et l'heure du jour : jour_mois_an_heure_min_sec
	anneeMoisJour := now.Format("01_02_2006_15_04_05")

	nomFichier := adrClient + "_" + anneeMoisJour + "_" + fichierId + ".csv"

	// Pour éviter des erreurs, on supprime les caractères spéciaux du nom du fichier
	c_a_remplacer := []string{"[", "]", ":"}
	for _, elt := range c_a_remplacer {
		nomFichier = strings.Replace(nomFichier, elt, "", -1)
	}

	return nomFichier
}

func ecriture_csv(data []byte, fichier string) {
	// Crée un nouveau fichier, si il existe, il est réinitialisé
	nvFichier, err := os.OpenFile(fichier, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	defer nvFichier.Close() // Ferme le fichier a la fin de la fonction
	gestion_erreur(err, "Creation du fichier")

	// Ecris les données dans le fichier créé
	err = os.WriteFile(fichier, data, 0644)
	gestion_erreur(err, "Ecriture dans fichier")
}

func main() {
	gob.Register([]byte{}) // Enregistre le type []byte pour le gob

	// On récupère les arguments de la ligne de commande
	arguments := os.Args
	if len(arguments) != 2 { // On sort du main si les arguments ne sont pas corrects
		fmt.Printf("Argument incorrect \n")
		fmt.Printf("go run programme.go <port>")
		return
	}
	port := arguments[1]

	// Connexion du serveur
	conn, err := net.Listen("tcp", ":"+port)
	gestion_erreur(err, "Serveur à l'écoute")
	defer conn.Close() // Ferme la connexion à la fin du programme

	for { // Boucle infinie qui permet d'accepter tous les clients
		client, err := conn.Accept()
		gestion_erreur(err, "Connexion du client")

		go func(client net.Conn) { // Permet de gérer plusieurs clients simultanément
			defer client.Close() // Ferme la connexion à la fin de la goroutine
			decoder := gob.NewDecoder(client)

			addrClient := client.RemoteAddr().String() // Récupère l'adresse ip du client

			// Récupère les paramètres
			var dataP []byte
			err = decoder.Decode(&dataP)
			gestion_erreur(err, "Décodage")
			param := string(dataP)
			listeParam := strings.Split(param, " ")
			nb_goroutines, err := strconv.Atoi(listeParam[0])
			gestion_erreur(err, "Conversion en int")
			dist_limite, err := strconv.Atoi(listeParam[1])
			gestion_erreur(err, "Conversion en int")

			// Récupère le premier fichier CSV
			nomFichier1 := nom_fichier(addrClient, "1")
			var data1 []byte
			err = decoder.Decode(&data1)
			gestion_erreur(err, "Décodage")
			ecriture_csv(data1, nomFichier1)

			// Récupère le second fichier CSV
			nomFichier2 := nom_fichier(addrClient, "2")
			var data2 []byte
			err = decoder.Decode(&data2)
			gestion_erreur(err, "Décodage")
			ecriture_csv(data2, nomFichier2)

			// Appel du main pour le traitement (à implémenter)
			fmt.Printf("Les paramètres récupérés sont : %v et %v", nb_goroutines, dist_limite)
			var data = []byte("Test")

			// On renvoie le fichier avec les distances de Levenshtein au client
			encoder := gob.NewEncoder(client)
			err = encoder.Encode(data)
			gestion_erreur(err, "Envoi via TCP")
		}(client)

	}

}
