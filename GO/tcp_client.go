package main

import (
	"fmt"
	"net"
	"os"
)

func bdd_to_binaire(chemin string) ([]byte, int) {
	file, err := os.Open(chemin)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier : %v\n", err)
		return nil, 0
	}
	// On récupère la taille du fichier en octets
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des infos du fichier : %v\n", err)
		return nil, 0
	}
	sizeFile := fileInfo.Size()

	data := make([]byte, sizeFile)
	n, err := file.Read(data)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return nil, 0
	}
	defer file.Close()

	// Retourne l'intégralité du fichier en binaire et le nombre d'octets
	return data[:n], int(sizeFile)
}

func envoi_segmentation(conn net.Conn, data []byte, segment int, tailleFichier int) { // Pas testé
	for i := 0; i < int(tailleFichier/segment); i++ {
		_, err := conn.Write(data[segment*i : segment*(i+1)])
		if err != nil {
			fmt.Printf("Erreur lors de l'envoi du message : %v\n", err)
			return
		}
	}
	// Message de fin d'envoi à BDD1
	msgFin := []byte("Fin envoi BDD1")
	_, err := conn.Write(msgFin)
	if err != nil {
		fmt.Printf("Erreur lors de l'envoi du message : %v\n", err)
		return
	}
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

	// Connexion du client
	conn, err := net.Dial("tcp", serv+":"+port)
	if err != nil {
		fmt.Printf("Erreur de connexion au serveur: %v\n", err)
		return
	}
	defer conn.Close() // Ferme la connexion à la fin du programme

	// On transforme les bases de données en binaire
	data1, nbOctet1 := bdd_to_binaire(bdd1)
	data2, nbOctet2 := bdd_to_binaire(bdd2)

	// On segmente la base de données pour éviter la perte de données (par paquet de 1024)
	// Pas testé
	taillePaquet := 1024
	envoi_segmentation(conn, data1, taillePaquet, nbOctet1)
	envoi_segmentation(conn, data2, taillePaquet, nbOctet2)

	// On envoie un message pour indiquer la fin du transfert au serveur

	msg := []byte("Hello Word")
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Printf("Erreur lors de l'envoi du message : %v\n", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Erreur lors de la réception du message : %v\n", err)
		return
	}

	fmt.Printf("Message reçu du serveur : %s\n", string(buffer[:n]))

}
