package main

import (
	"fmt"
	"net"
	"os"
)

func lecture_msg(client net.Conn) {
	buffer := make([]byte, 1024)
	n, err := client.Read(buffer)
	if err != nil {
		fmt.Printf("Erreur lors de la réception du message : %v\n", err)
		return
	}
	fmt.Printf("Message reçu : %s\n", string(buffer[:n]))

	msg := []byte("Bien reçu")
	_, err = client.Write(msg)
	if err != nil {
		fmt.Printf("Erreur lors de l'envoi du message : %v\n", err)
		return
	}
	defer client.Close()

}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("Argument manquant \n")
		fmt.Printf("Entrez le numéro de port\n")
		return
	}
	if len(arguments) > 2 {
		fmt.Printf("Trop d'arguments\n")
		fmt.Printf("Entrez le numéro de port\n")
		return
	}

	port := arguments[1]
	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Erreur de connexion au serveur: %v\n", err)
		return
	}
	defer conn.Close()

	for {
		client, err := conn.Accept()
		if err != nil {
			fmt.Printf("Erreur lors de la connexion avec le client : %v", err)
			return
		}
		go lecture_msg(client)

	}

}
