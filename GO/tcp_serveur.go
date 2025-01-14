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

func bytes_to_csv(data []byte, fichier string) {
	nvFichier, err := os.OpenFile(fichier, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	defer nvFichier.Close()
	gestion_erreur(err, "Creation du fichier")
	err = os.WriteFile(fichier, data, 0644)
	gestion_erreur(err, "Ecriture dans fichier")
}

func lecture_msg(client net.Conn) {
	defer client.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := client.Read(buffer)
		gestion_erreur(err, "Reception via TCP")
		if string(buffer[:n]) == "WAIT" {
			fmt.Printf("On va envoyer des fichiers\n")
			break
		} else {

			// Premier fichier
			if string(buffer[:n]) == "BEGIN" {
				data1, nom1 := resegmentation(client)
				bytes_to_csv(data1, nom1)
			}

			// Second fichier
			if string(buffer[:n]) == "BEGIN" {
				data2, nom2 := resegmentation(client)
				bytes_to_csv(data2, nom2)
			}
		}
	}
}

func resegmentation(client net.Conn) ([]byte, string) {
	buffer := make([]byte, 1024)
	n, err := client.Read(buffer)
	gestion_erreur(err, "Reception via TCP")

	// Lecture de l'en-tête
	entete := strings.Split(string(buffer[:n]), " ")
	fmt.Printf("%q", entete[:])
	if len(entete) != 3 {
		fmt.Println("Entête invalide reçue")
		os.Exit(1) // Arrête le programme avec un code de retour 1 (indiquant une erreur)
	}
	nomFichier := entete[0]
	tailleFichier, err := strconv.Atoi(entete[1])
	gestion_erreur(err, "Conversion string vers int")

	taillePaquet, err := strconv.Atoi(entete[2])
	gestion_erreur(err, "Conversion string vers int")

	data := make([]byte, 0, tailleFichier)

	for {
		buffer = make([]byte, int(taillePaquet))
		n, err = client.Read(buffer)
		gestion_erreur(err, "Réception via TCP")
		data = append(data, buffer[:n]...)
		if string(buffer[:n]) != "END" {
			break
		}
	}
	return data, nomFichier
}

func main() {
	// On récupère les arguments de la ligne de commande
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Printf("Argument incorrect \n")
		fmt.Printf("go run programme.go <port>")
		return
	}
	port := arguments[1]

	// Connexion du client
	conn, err := net.Listen("tcp", ":"+port)
	gestion_erreur(err, "Serveur à l'écoute")
	defer conn.Close()

	for {
		client, err := conn.Accept()
		gestion_erreur(err, "Connexion du client")
		go lecture_msg(client)

	}

}
