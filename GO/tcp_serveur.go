package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func gestion_erreur(err error, message string) {
	if err != nil {
		fmt.Printf("Erreur : %s - %v\n", message, err)
		os.Exit(1) // Arrête le programme avec un code de retour 1 (indiquant une erreur)
	}
}

func nom_fichier(fichierId string, adrClient string) string {
	// Date
	now := time.Now()
	anneeMoisJour := now.Format("2006_01_02")
	// Nom du fichier
	nomFichier := adrClient + "_" + anneeMoisJour + "_" + fichierId + ".csv"
	c_a_remplacer := []string{"[", "]", ":"}
	for _, elt := range c_a_remplacer {
		nomFichier = strings.Replace(nomFichier, elt, "", -1)
	}

	return nomFichier
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
		go func(client net.Conn) {
			defer client.Close()
			decoder := gob.NewDecoder(client)

			addrClient := client.RemoteAddr().String()
			nomFichier1 := nom_fichier(addrClient, "1")
			var data1 []byte
			err := decoder.Decode(&data1)
			gestion_erreur(err, "Décodage")
			ecriture_csv(data1, nomFichier1)

			nomFichier2 := nom_fichier(addrClient, "2")
			var data2 []byte
			err = decoder.Decode(&data2)
			gestion_erreur(err, "Décodage")
			ecriture_csv(data2, nomFichier2)

			// Appel du main
			var data = []byte("Test")

			// On renvoie le fichier avec les distances de Levenshtein au client
			encoder := gob.NewEncoder(client)
			err = encoder.Encode(data)
			gestion_erreur(err, "Envoi via TCP")
		}(client)

	}

}
