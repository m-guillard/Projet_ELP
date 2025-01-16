package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func gestion_erreur(err error, message string) {
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

	// Retourne l'intégralité du fichier en binaire et le nombre d'octets
	return data
}

func ecriture_csv(data []byte, fichier string) {
	nvFichier, err := os.OpenFile(fichier, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	defer nvFichier.Close()
	gestion_erreur(err, "Creation du fichier")
	err = os.WriteFile(fichier, data, 0644)
	gestion_erreur(err, "Ecriture dans fichier")
}

func main() {
	gob.Register([]byte{})
	// On récupère les arguments de la ligne de commande
	arguments := os.Args
	if len(arguments) != 5 {
		fmt.Printf("Argument incorrect \n")
		fmt.Printf("go run programme.go <adresse serveur> <port> <chemin base de donnee 1> <chemin base de donnee 2>")
		return
	}
	serv, port, bdd1, bdd2 := arguments[1], arguments[2], arguments[3], arguments[4]

	// Connexion du client
	conn, err := net.Dial("tcp", serv+":"+port)
	gestion_erreur(err, "Connexion au serveur")
	defer conn.Close() // Ferme la connexion à la fin du programme

	// On transforme les bases de données en binaire
	data1 := lecture_csv(bdd1)
	data2 := lecture_csv(bdd2)

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	// Envoi les données au serveur
	err = encoder.Encode(data1)
	gestion_erreur(err, "Envoi via TCP")
	err = encoder.Encode(data2)
	gestion_erreur(err, "Envoi via TCP")

	fmt.Print("Fichiers envoyés au serveur\n")

	// Réception du fichier final
	var data []byte
	nFichier := "final.csv"
	err = decoder.Decode(&data)
	gestion_erreur(err, "Décodage")
	ecriture_csv(data, nFichier)
	fmt.Printf("Fin de traitement, %q reçu\n", nFichier)

}
