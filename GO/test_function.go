package main

import (
	"fmt"
	"sync"
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
