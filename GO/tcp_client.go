package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func gestion_erreur(err error, message string) {
	// Arrête le programme si il y a une erreur et affiche un message d'erreur dans le terminal
	if err != nil {
		fmt.Printf("Erreur : %s - %v\n", message, err)
		os.Exit(1) // Arrête le programme avec un code de retour 1 (indiquant une erreur)
	}
}

func lecture_csv(chemin string) []byte {
	// Ouverture du fichier CSV à lire
	file, err := os.Open(chemin)
	gestion_erreur(err, "Ouverture fichier")
	defer file.Close()

	// On récupère la taille du fichier en octets
	fileInfo, err := file.Stat()
	gestion_erreur(err, "Statistiques du fichier")
	sizeFile := fileInfo.Size()

	// Lecture des données
	data := make([]byte, sizeFile)
	_, err = file.Read(data)
	gestion_erreur(err, "Lecture du fichier")
	defer file.Close()

	// Retourne l'intégralité du fichier en binaire
	return data
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
	if len(arguments) != 7 { // On sort du main si les arguments ne sont pas corrects
		fmt.Printf("Argument incorrect \n")
		fmt.Printf("go run programme.go <adresse serveur> <port> <chemin base de donnee 1> <chemin base de donnee 2> <nombre goroutines> <distance Levenshtein limite>")
		return
	}
	serv, port, bdd1, bdd2, nb_goroutines, dist_limite := arguments[1], arguments[2], arguments[3], arguments[4], arguments[5], arguments[6]

	// Connexion du client
	conn, err := net.Dial("tcp", serv+":"+port)
	gestion_erreur(err, "Connexion au serveur")
	defer conn.Close() // Ferme la connexion à la fin du programme

	// On transforme les bases de données en binaire
	data1 := lecture_csv(bdd1)
	data2 := lecture_csv(bdd2)

	// On transforme les paramètres en binaire
	dataP := []byte(nb_goroutines + " " + dist_limite)

	// Envoi les données au serveur
	encoder := gob.NewEncoder(conn)
	// Envoie les paramètres au serveur
	err = encoder.Encode(dataP)
	gestion_erreur(err, "Envoi via TCP")
	// Envoie la première base de données
	err = encoder.Encode(data1)
	gestion_erreur(err, "Envoi via TCP")
	// Envoie la seconde base de données
	err = encoder.Encode(data2)
	gestion_erreur(err, "Envoi via TCP")

	fmt.Print("Fichiers envoyés au serveur\n")

	// Réception du fichier final de la part du serveur
	decoder := gob.NewDecoder(conn)
	var data []byte
	nFichier := "final.csv"
	err = decoder.Decode(&data)
	gestion_erreur(err, "Décodage")
	ecriture_csv(data, nFichier)
	fmt.Printf("Fin de traitement, %q reçu\n", nFichier)
}
