package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Printf("Argument manquant \n")
		fmt.Printf("Entrez le serveur et un numéro de port\n")
		return
	}
	if len(arguments) > 3 {
		fmt.Printf("Trop d'arguments\n")
		fmt.Printf("Entrez le serveur et un numéro de port\n")
		return
	}

	serv, port := arguments[1], arguments[2]
	conn, err := net.Dial("tcp", serv+":"+port)
	if err != nil {
		fmt.Printf("Erreur de connexion au serveur\n")
		return
	}
	defer conn.Close()

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
