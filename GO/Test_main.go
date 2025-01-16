// V2

// Calcul parallèle des distances
func deroule(noms1, noms2 []string, dist_max, numGoRoutines int, safeMap *SafeMap) {
	var wg sync.WaitGroup
	tasks := make(chan [2]string, len(noms1)*len(noms2))

	go func() {
		for _, nomA := range noms1 {
			for _, nomB := range noms2 {
				tasks <- [2]string{nomA, nomB}
			}
		}
		close(tasks)
	}()

	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pair := range tasks {
				nomA, nomB := pair[0], pair[1]
				mat := matrice(nomA, nomB)
				distance := matrice_lev(nomA, nomB, mat)
				safeMap.MapLevenshtein(nomA, nomB, distance, dist_max)
			}
		}()
	}

	wg.Wait()
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage : go run main.go <fichier1.csv> <colonne1> <fichier2.csv> <colonne2> [dist_max] [nb_goroutines]")
		return
	}

	fichier1 := os.Args[1]
	colonne1 := os.Args[2]
	fichier2 := os.Args[3]
	colonne2 := os.Args[4]
	//go run main.go fichier1.csv fichier2.csv colonne1 colonne2 3 4
	// os.Args[0] : Contient le nom de l'exécutable ou de la commande (main.go ici). On l'ignore généralement pour la récupération des paramètres.
	// os.Args[1] : Premier fichier CSV (fichier1.csv).
	// os.Args[2] : Second fichier CSV (fichier2.csv).
	// os.Args[3] : Colonne cible dans le premier fichier (colonne1).
	// os.Args[4] : Colonne cible dans le second fichier (colonne2).
	// os.Args[5] : Distance maximale de Levenshtein (3).
	// os.Args[6] : Nombre de goroutines (4).

	dist_max := 3
	// un 6ᵉ argument (index 5) est fourni pour dist_max ? Si oui, lu et assigné. 
	if len(os.Args) > 5 {
		// fmt.Sscanf ==> convertir un argument en entier de manière sûre
		fmt.Sscanf(os.Args[5], "%d", &dist_max)
	}

	nb_goroutines := 4
	if len(os.Args) > 6 {
		fmt.Sscanf(os.Args[6], "%d", &nb_goroutines)
	}

	noms1 := extractionColonne(fichier1, colonne1)
	
	noms2 := extractionColonne(fichier2, colonne2)
	

	c := SafeMap{map_lev: make(map[string]map[string]int)}
	start := time.Now()
	deroule(noms1, noms2, dist_max, nb_goroutines, &c)
	fmt.Printf("Temps de calcul : %v\n", time.Since(start))
	}





// V1

// Calcul parallèle des distances
func deroule(names1, names2 []string, dist_max, numGoRoutines int, safeMap *SafeMap) {
	var wg sync.WaitGroup
	tasks := make(chan [2]string, len(names1)*len(names2))

	go func() {
		for _, nameA := range names1 {
			for _, nameB := range names2 {
				tasks <- [2]string{nameA, nameB}
			}
		}
		close(tasks)
	}()

	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pair := range tasks {
				nameA, nameB := pair[0], pair[1]
				mat := matrice(nameA, nameB)
				distance := matrice_lev(nameA, nameB, mat)
				safeMap.MapLevenshtein(nameA, nameB, distance, dist_max)
			}
		}()
	}

	wg.Wait()
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage : go run main.go <fichier1.csv> <colonne1> <fichier2.csv> <colonne2> [dist_max] [nb_goroutines]")
		return
	}

	fichier1 := os.Args[1]
	colonne1 := os.Args[2]
	fichier2 := os.Args[3]
	colonne2 := os.Args[4]

	dist_max := 3
	if len(os.Args) > 5 {
		fmt.Sscanf(os.Args[5], "%d", &dist_max)
	}

	nb_goroutines := 4
	if len(os.Args) > 6 {
		fmt.Sscanf(os.Args[6], "%d", &nb_goroutines)
	}

	names1, err := extractionColonne(fichier1, colonne1)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	names2, err := extractionColonne(fichier2, colonne2)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	c := SafeMap{map_lev: make(map[string]map[string]int)}
	start := time.Now()
	deroule(fichier1, fichier2, dist_max, nb_goroutines, &c)
	fmt.Printf("Temps de calcul : %v\n", time.Since(start))

}


// V0

func deroule() {
	// On a deux mots
	var mot_A string = "CHAT"
	var mot_B string = "CHIEN"
	dist_max := 1

	// On crée une matrice vide de longueur adaptée à ces deux mots
	matrice_vide := matrice(mot_A, mot_B)

	// On remplit cette matrice avec l'algo de Levenshtein et on retourne la valeur de la distance de Levenshtein
	Distance_Levenshtein := matrice_lev(mot_A, mot_B, matrice_vide)

	// On ajoute ce lien entre les mots au dictionnaire partagé (si la distance est inférieure à dist_max)
	MapLevenshtein(mot_A, mot_B, Distance_Levenshtein, dist_max)
}

func main() {
	
	
	// créer dico partagé, accès
	var wg sync.WaitGroup                                  // crée groupe d'attente, on y ajoute les go routines, pour attendre qu'elles finissent pour créer le csv
	c := SafeMap{map_lev: make(map[string]map[string]int)} // Création du channel
	
	// Attention, il faudra pouvoir faire un os.args à la place de ça
	fichier1 = csv.load
	fichier 2 = 

	// après, avec go routines, on aura pleins de mots à comparer
	// for qui lance les go routines, de la syntaxe suivante
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.MapLevenshtein("dinosaure", "brebis", 12)
	}()

	// récupération des go routines, transfo
	wg.Wait() // attend que toutes les goroutines soient terminées

	// param nomenclature fichier csv final
	now := time.Now()
	anneeMoisJour := now.Format("2006_01_02")

	dico_to_csv(c.map_lev, anneeMoisJour)

}