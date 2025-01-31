package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// Fonction pour obtenir la valeur minimale parmi 3 entiers
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func dico_to_csv(map_lev map[string]map[string]int, date string) string {
	// Construction nom fichier avec la date
	fileName := fmt.Sprintf("output_%s.csv", date)

	// Création fichier CSV
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return ""
	}
	defer file.Close()

	// Créer un writer CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Écrire l'en-tête
	header := []string{"Nom_A", "Nom_B", "Distance_Levenshtein"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Erreur lors de l'écriture de l'en-tête :", err)
		return ""
	}

	// Parcourir la map imbriquée et écrire les lignes
	for Nom_A, Nom_B_lie := range map_lev {
		for Nom_B, Distance_Levenshtein := range Nom_B_lie {
			row := []string{Nom_A, Nom_B, fmt.Sprintf("%d", Distance_Levenshtein)}
			if err := writer.Write(row); err != nil {
				fmt.Println("Erreur lors de l'écriture d'une ligne :", err)
				return ""
			}
		}
	}

	fmt.Println("Fichier CSV généré avec succès : output.csv")
	return fileName

}

// Structure partagée
type SafeMap struct {
	mu sync.Mutex
	// La map a pour clé le nom de la première base de données
	// Sa valeur est une map qui a pour clé le nom de la seconde base de données et comme valeur la distance de Levenshtein
	map_lev map[string]map[string]int
}

// Extrait une colonne d'un fichier CSV et l'enregistre dans un nouveau fichier CSV
// nomFichier : chemin du fichier où il faut extraire une colonne
// nomColonne : le nom de colonne à extraire
// extractionColonne ne change pas, elle continue de créer les fichiers CSV intermédiaires
func extractionColonne(nomFichier string, nomColonne string) string {
	// Lecture de nomFichier
	fichierOriginal, err := os.Open(nomFichier)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier : %v\n", err)
		return "" // Retourne une chaîne vide si erreur
	}
	defer fichierOriginal.Close()

	// Nom du fichier de sortie
	nomNvFichier := strings.Split(nomFichier, ".")[0] + "_" + nomColonne + ".csv" // Nom du nouveau fichier

	// Crée un fichier de sortie
	nvFichier, err := os.OpenFile(nomNvFichier, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier de sortie : %v\n", err)
		return "" // Retourne une chaîne vide si erreur
	}
	defer nvFichier.Close()

	// Scanner pour lire ligne par ligne
	scanner := bufio.NewScanner(fichierOriginal)
	var indiceColonne int = -1
	for scanner.Scan() {
		ligne := strings.Split(scanner.Text(), ";")
		if indiceColonne == -1 { // Recherche de l'indice correspondant au nom de colonne
			for index, elt := range ligne {
				if elt == nomColonne {
					indiceColonne = index
					break
				}
			}
			if indiceColonne == -1 {
				fmt.Printf("Nom de colonne non trouvé dans le fichier\n")
				return ""
			}
		} else {
			// Vérifie que la ligne est complète
			if indiceColonne < len(ligne) {
				// Écriture de la cellule correspondante dans le fichier intermédiaire
				_, err := nvFichier.WriteString(ligne[indiceColonne] + "\n")
				if err != nil {
					fmt.Printf("Erreur d'écriture dans le fichier : %v\n", err)
					return ""
				}
			} else {
				// Si la colonne n'existe pas dans la ligne, écrire une chaîne vide
				_, err := nvFichier.WriteString("\n")
				if err != nil {
					fmt.Printf("Erreur d'écriture dans le fichier : %v\n", err)
					return ""
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return ""
	}

	return nomNvFichier // Retourne le nom du fichier intermédiaire créé
}

// Met à jour la MapLevenshtein qui est une structure partagée
// Paramètres : motA et motB : les mots comparés ; dist_int : distance de Levenshtein entre motA et motB
// dist_max : si dist est inférieure à dist_max, on juge que motA et motB sont semblables et on conserve la valeur
func (s *SafeMap) MapLevenshtein(motA string, motB string, dist int, dist_max int) {
	if dist <= dist_max { // Ne pas prendre en compte les données si distance de Levenshtein trop élevée
		// Accès à la structure partagée
		s.mu.Lock()
		defer s.mu.Unlock()
		valeur, existe := s.map_lev[motA] // Récupère la map associé à la clé A

		if existe == false {
			valeur = make(map[string]int)
		}
		valeur[motB] = dist // On ajoute dans le map, le nouveau mot avec sa distance de Levenshtein
		s.map_lev[motA] = valeur
	}
}

// Affiche la Map, utile pour le deboggage
func (s *SafeMap) Display() {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("\nDictionnaire:\n")
	for k, v := range s.map_lev {
		fmt.Printf("%v \n", k)
		for k2, v2 := range v {
			fmt.Printf("-> %v = %v \n", k2, v2)
		}
	}
}

// Initialise la matrice pour l'algorithme de Leveinstein
func matrice(motA string, motB string) [][]int {

	var longMotA, longMotB int = len(motA) + 1, len(motB) + 1

	mat := make([][]int, longMotA) // Le nombre de ligne correspond à la longueur du mot A+1
	for i := range mat {           // Le nombre de colonne correspond à la longueur du mot B+1
		mat[i] = make([]int, longMotB)
	}

	// On remplit la matrice qui est à 0
	for i := 0; i < longMotA; i++ {
		// Première colonne rempli de 0 à la longueur du mot A
		mat[i][0] = i
	}
	for i := 0; i < longMotB; i++ {
		// Première ligne rempli de 0 à la longueur du mot B
		mat[0][i] = i
	}

	return mat
}

func matrice_lev(mot_A, mot_B string, matrice [][]int) int {
	liste_A := strings.Split(mot_A, "")
	//fmt.Println(liste_A)
	liste_B := strings.Split(mot_B, "")
	//fmt.Println(liste_B)

	for i := 1; i < len(liste_A)+1; i++ {
		for j := 1; j < len(liste_B)+1; j++ {
			cout_substitution := 1
			// si la lettre i du mot A est identique à la lettre j du mot B, alors la valeur de la cellule est 0
			if liste_A[i-1] == liste_B[j-1] {
				cout_substitution = 0
			}
			// Calcul des coûts : insertion, suppression, substitution
			insertion := matrice[i][j-1] + 1
			suppression := matrice[i-1][j] + 1
			substitution := matrice[i-1][j-1] + cout_substitution
			matrice[i][j] = min(insertion, suppression, substitution)

		}
	}

	// Affichage de la matrice remplie
	// fmt.Println("Matrice remplie :")
	// for _, row := range matrice_vide {
	// 	for _, value := range row {
	// 		fmt.Printf("%d ", value)
	// 	}
	// 	fmt.Println()
	// }
	return matrice[len(liste_A)][len(liste_B)]
}

// Fonction qui compare les mots en lisant depuis les fichiers CSV
func derouleDepuisCSV(fichier1, fichier2 string, dist_max, numGoRoutines int, safeMap *SafeMap) {
	var wg sync.WaitGroup
	tasks := make(chan [2]string, 1000) // Taille du canal à ajuster selon la mémoire

	// Ouverture des fichiers CSV
	f1, err := os.Open(fichier1)
	if err != nil {
		fmt.Printf("Erreur d'ouverture du fichier %s: %v\n", fichier1, err)
		return
	}
	defer f1.Close()

	f2, err := os.Open(fichier2)
	if err != nil {
		fmt.Printf("Erreur d'ouverture du fichier %s: %v\n", fichier2, err)
		return
	}
	defer f2.Close()

	// Scanner pour lire les fichiers ligne par ligne
	scanner1 := bufio.NewScanner(f1)

	// Lancer une goroutine pour ajouter des tâches dans le canal
	go func() {
		for scanner1.Scan() {
			nomA := scanner1.Text()

			// Ignorer les lignes vides dans fichier1
			if strings.TrimSpace(nomA) == "" {
				continue // Passe à la ligne suivante si la ligne est vide
			}

			// Revenir au début du fichier 2 pour chaque ligne de fichier 1
			f2.Seek(0, 0) // Rewind du fichier 2
			scanner2 := bufio.NewScanner(f2)

			// Comparer nomA avec toutes les lignes de fichier 2
			for scanner2.Scan() {
				nomB := scanner2.Text()

				// Ignorer les lignes vides dans fichier2
				if strings.TrimSpace(nomB) == "" {
					continue // Passe à la ligne suivante si la ligne est vide
				}

				tasks <- [2]string{nomA, nomB} // Ajouter la paire de mots dans le canal
			}
		}
		close(tasks) // Fermer le canal une fois toutes les tâches ajoutées
	}()

	// Lancer les goroutines pour traiter les paires de mots
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pair := range tasks {
				nomA, nomB := pair[0], pair[1]
				mat := matrice(nomA, nomB)                             // Matrice de Levenshtein
				distance := matrice_lev(nomA, nomB, mat)               // Calcul de la distance
				safeMap.MapLevenshtein(nomA, nomB, distance, dist_max) // Stockage dans la map sécurisée
			}
		}()
	}

	wg.Wait() // Attente de la fin de toutes les goroutines
}

func Main_code(fichier1 string, colonne1 string, fichier2 string, colonne2 string, dist_max int, nb_goroutines int) string {
	// if len(os.Args) < 5 {
	//	fmt.Println("Usage : go run main.go <fichier1.csv> <colonne1> <fichier2.csv> <colonne2> [dist_max] [nb_goroutines]")
	// return
	// }

	// fichier1 := os.Args[1]
	// colonne1 := os.Args[2]
	// fichier2 := os.Args[3]
	// colonne2 := os.Args[4]

	// dist_max := 3
	// un 6ᵉ argument (index 5) est fourni pour dist_max ? Si oui, lu et assigné.
	if len(os.Args) > 5 {
		// fmt.Sscanf ==> convertir un argument en entier de manière sûre
		fmt.Sscanf(os.Args[5], "%d", &dist_max)
	}

	// nb_goroutines := 4
	// if len(os.Args) > 6 {
	// 	fmt.Sscanf(os.Args[6], "%d", &nb_goroutines)
	// }

	// Récupérer les colonnes extraites sous forme de slices
	noms1 := extractionColonne(fichier1, colonne1)
	noms2 := extractionColonne(fichier2, colonne2)

	// Crée la structure SafeMap
	c := SafeMap{map_lev: make(map[string]map[string]int)}
	start := time.Now()

	// Appel à deroule avec les noms extraits
	// Appel à la fonction derouleDepuisCSV pour traiter les fichiers intermédiaires
	derouleDepuisCSV(noms1, noms2, dist_max, nb_goroutines, &c)
	fmt.Printf("Temps de calcul : %v\n", time.Since(start))

	// Création du fichier CSV final
	now := time.Now()
	anneeMoisJour := now.Format("2006_01_02")
	fileName := dico_to_csv(c.map_lev, anneeMoisJour)

	return fileName
}
