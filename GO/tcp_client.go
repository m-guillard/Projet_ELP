package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func gestion_erreur(err error, message string) {
	if err != nil {
		fmt.Printf("Erreur : %s - %v\n", message, err)
		os.Exit(1) // Arrête le programme avec un code de retour 1 (indiquant une erreur)
	}
}

func bdd_to_binaire(chemin string) ([]byte, int) {
	file, err := os.Open(chemin)
	gestion_erreur(err, "Ouverture fichier")
	// On récupère la taille du fichier en octets
	fileInfo, err := file.Stat()
	gestion_erreur(err, "Statistiques du fichier")
	sizeFile := fileInfo.Size()

	data := make([]byte, sizeFile)
	n, err := file.Read(data)
	gestion_erreur(err, "Lecture du fichier")
	defer file.Close()

	// Retourne l'intégralité du fichier en binaire et le nombre d'octets
	return data[:n], int(sizeFile)
}

func envoi_segmentation(conn net.Conn, data []byte, segment int, tailleFichier int) { // Pas testé
	for i := 0; i < int(tailleFichier/segment); i++ {
		_, err := conn.Write(data[segment*i : segment*(i+1)])
		gestion_erreur(err, "Envoi via TCP")
	}
	reste := tailleFichier % segment
	if reste > 0 {
		_, err := conn.Write(data[tailleFichier-reste:])
		gestion_erreur(err, "Envoi via TCP des données restantes")
	}
	// Message de fin d'envoi à BDD
	msgFin := []byte("END")
	_, err := conn.Write(msgFin)
	gestion_erreur(err, "Envoi via TCP")
}

func entete(conn net.Conn, nomFichier string, tailleFichier int, taillePaquet int) {

	_, err := conn.Write([]byte("BEGIN"))
	gestion_erreur(err, "Envoi via TCP")

	msg := []byte(nomFichier + " " + strconv.Itoa(tailleFichier) + " " + strconv.Itoa(taillePaquet))
	fmt.Printf("%q\n", msg)
	_, err = conn.Write(msg)
	gestion_erreur(err, "Envoi via TCP")
}

func main() {
	// On récupère les arguments de la ligne de commande
	arguments := os.Args
	if len(arguments) != 5 {
		fmt.Printf("Argument incorrect \n")
		fmt.Printf("go run programme.go <adresse serveur> <port> <chemin base de donnee 1> <chemin base de donnee 2>")
		return
	}
	serv, port, bdd1, bdd2 := arguments[1], arguments[2], arguments[3], arguments[4]
	splitPath1 := strings.Split(bdd1, "/")
	nom1 := splitPath1[len(splitPath1)-1]

	splitPath2 := strings.Split(bdd2, "/")
	nom2 := splitPath2[len(splitPath2)-1]

	// Connexion du client
	conn, err := net.Dial("tcp", serv+":"+port)
	gestion_erreur(err, "Connexion au serveur")
	defer conn.Close() // Ferme la connexion à la fin du programme

	// On transforme les bases de données en binaire
	data1, nbOctet1 := bdd_to_binaire(bdd1)
	data2, nbOctet2 := bdd_to_binaire(bdd2)

	// On segmente la base de données pour éviter la perte de données (par paquet de 1024)
	// Pas testé
	taillePaquet := 1024
	entete(conn, nom1, nbOctet1, taillePaquet)
	envoi_segmentation(conn, data1, taillePaquet, nbOctet1)
	entete(conn, nom2, nbOctet2, taillePaquet)
	envoi_segmentation(conn, data2, taillePaquet, nbOctet2)

	// On attend que le serveur réponde par un fichier
	msg := []byte("WAIT")
	_, err = conn.Write(msg)
	gestion_erreur(err, "Envoi via TCP")

}
