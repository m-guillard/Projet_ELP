package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

func dico_to_csv(map_lev map[string]map[string]int, date string) {
	// Construction nom fichier avec la date
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
