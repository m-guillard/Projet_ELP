package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// Structure partagée
type SafeMap struct {
	mu      sync.Mutex
	map_lev map[string]map[string]int
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

// Renvoie un map mis à jour. Le map a pour clé le nom de la première base de données
// Sa valeur est une map qui a pour clé le nom de la seconde base de données et comme valeur la distance de Levenshtein
func (s *SafeMap) MapLevenshtein(motA string, motB string, dist int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	valeur, existe := s.map_lev[motA] // Récupère le map associé à la clé A
	// A faire : GESTION DE DOUBLONS
	if existe == false {
		valeur = make(map[string]int)
	}
	valeur[motB] = dist // On ajoute dans le map, le nouveau mot avec sa distance de Levenshtein
	s.map_lev[motA] = valeur
}

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

func dico_to_csv(map_lev map[string]map[string]int, date string) {
	// Construction nom fichier avec la datee
	fileName := fmt.Sprintf("output_%s.csv", date)

	// Création fichier CSV
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}
	defer file.Close()

	// Créer un writer CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Écrire l'en-tête
	header := []string{"Nom_A", "Nom_B", "Distance_Levenshtein"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Erreur lors de l'écriture de l'en-tête :", err)
		return
	}

	// Parcourir la map imbriquée et écrire les lignes
	for Nom_A, Nom_B_lie := range map_lev {
		for Nom_B, Distance_Levenshtein := range Nom_B_lie {
			row := []string{Nom_A, Nom_B, fmt.Sprintf("%d", Distance_Levenshtein)}
			if err := writer.Write(row); err != nil {
				fmt.Println("Erreur lors de l'écriture d'une ligne :", err)
				return
			}
		}
	}

	fmt.Println("Fichier CSV généré avec succès : output.csv")

}

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

func comparaison(mot_A, mot_B string, matrice_vide [][]int) {
	liste_A := strings.Split(mot_A, "")
	fmt.Println(liste_A)
	liste_B := strings.Split(mot_B, "")
	fmt.Println(liste_B)

	for i := 1; i < len(liste_A)+1; i++ {
		for j := 1; j < len(liste_B)+1; j++ {
			cout_substitution := 1
			// si la lettre i du mot A est identique à la lettre j du mot B, alors la valeur de la cellule est 0
			if liste_A[i-1] == liste_B[j-1] {
				cout_substitution = 0
			}
			// Calcul des coûts : insertion, suppression, substitution
			insertion := matrice_vide[i][j-1] + 1
			suppression := matrice_vide[i-1][j] + 1
			substitution := matrice_vide[i-1][j-1] + cout_substitution
			matrice_vide[i][j] = min(insertion, suppression, substitution)

		}
	}

	// Affichage de la matrice remplie
	fmt.Println("Matrice remplie :")
	for _, row := range matrice_vide {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

func main() {
	mot_A := "ALICE"
	mot_B := "MARION"

	// Matrice vide donnée par un autre code
	matrice_vide := [][]int{
		{0, 1, 2, 3, 4, 5, 6},
		{1, 0, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 0},
		{4, 0, 0, 0, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0},
	}

	map_lev := map[string]map[string]int{
		"nom": {
			"nom_2": 2,
			"nom_3": 3,
		},
		"nom_autre": {
			"nom_2": 20,
			"nom_4": 40,
		},
	}

	now := time.Now()
	anneeMoisJour := now.Format("2006_01_02")

	comparaison(mot_A, mot_B, matrice_vide)
	dico_to_csv(map_lev, anneeMoisJour)
}

func main() {
	var motA string = "CHAT"
	var motB string = "CHIEN"
	var mat [][]int = matrice(motA, motB)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			fmt.Printf("%v", mat[i][j])
		}
		fmt.Printf("\n")
	}
	var wg sync.WaitGroup

	// Partie avec la goroutine
	c := SafeMap{map_lev: make(map[string]map[string]int)} // Création du channel
	sous_ex2 := make(map[string]int)
	sous_ex2["chat"] = 1
	c.map_lev["chien"] = sous_ex2
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.MapLevenshtein("chien", "niche", 5)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.MapLevenshtein("dinosaure", "brebis", 12)
	}()
	wg.Wait()
	c.Display()

}

// à faire : main avec tout le code, serveur tcp
