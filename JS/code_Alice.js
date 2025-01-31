const prompt = require('prompt-sync')();  // Charger le module prompt-sync

function manche(nb_joueur, mot_a_deviner){
    console.log(`\nDébut de la maanche.`);
    indices(nb_joueur, mot_a_deviner);  // Collecte des indices
}

function jeu(){
    let nb_joueur = prompt('Nombre de joueurs ? ');
    for (let num_manche = 1; num_manche <= 4; num_manche++){
        manche(nb_joueur)
    }
}

function comparaison(mot_a, mot_b){
    // True si mots pareils, ou racines identiques
    if (mot_a === mot_b){
        return True
    }
    //else {
    //    for (let iteration = 0; iteration <mot_a.lenght; iteration++){
    //        // à voir comment on définit la racine
    //        pass
    //    }
    //} 
    else {
        return False
    }
    
}
function indices(nb_joueur, mot_a_deviner) {
    let liste_indice = [];  // Liste où on mettra les indices à chaque manche

    for (let joueur = 0; joueur < nb_joueur; joueur++) {
        let indice = prompt(`Joueur ${joueur + 1}, entre ton indice : `);

        // Vérifier que l'indice n'est pas trop similaire au mot à deviner
        while (comparaison(mot_a_deviner, indice)) {
            console.log('Joueur ${joueur + 1}, entre un autre indice, celui-ci est trop similaire au mot à deviner : ');
            indice = prompt(`Joueur ${joueur + 1}, entre un autre indice : `);
        }

        // Check similaritudes avec les indices précédents. Si pas similaire, ajoute indice à la liste
        let estUnique = true;
        for (let i = 0; i < liste_indice.length; i++) {
            if (comparaison(liste_indice[i], indice)) {
                estUnique = false;
                console.log('Cet indice a déjà été donné. Entrer un autre indice.');
                break;
            }
        }

        // Ajouter l'indice si il est unique
        if (estUnique) {
            liste_indice.push(indice);
        } else {
            joueur--;  // Redemander un indice pour ce joueur, incrémentation à l'envers
        }
    }

    // Afficher tous les indices collectés
    console.log("\nListe des indices des joueurs :");
    console.log(liste_indice);
}

jeu();