const prompt = require('prompt-sync')();  // Charger le module prompt-sync

function deroule(nb_joueur) {
    let liste_indice = []; // Liste où on

    for (let joueur = 0; joueur < nb_joueur; joueur++) {
        // Demander à chaque joueur de donner un indice (input)
        const indice = prompt(`Joueur ${joueur + 1}, entre ton indice : `);
        
        // Ajouter l'indice à la liste des indices
        liste_indice.push(indice);
    }

    // Afficher tous les indices collectés
    console.log("Liste des indices des joueurs :");
    console.log(liste_indice);
}

// Appeler la fonction avec un exemple (par exemple 3 joueurs)
deroule(3);
