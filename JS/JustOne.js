// PS C:\Users\alice\Documents\Alice\INSA\3A\S1\Projets_2\ELP\Projet_ELP\JS> node code_Alice.js

const prompt = require('prompt-sync')();  // Charger le module prompt-sync
const fs = require('fs');  // Correct


function manche(nb_joueur, mot_a_deviner){
    console.log(`\nDébut de la maanche.`);
    indices(nb_joueur, mot_a_deviner);  // Collecte des indices
}


function generation_pioche(nombre_mots, fichier){
    let mots = fs.readFileSync(fichier, 'utf8').split('\n').map(mot => mot.trim()); // Correction de FileSystem -> fs
    const liste_mots_deviner = [];
    for (let i = 0; i < nombre_mots; i++){
        let mot_alea = mots[Math.floor(Math.random() * mots.length)]; // Ajout de mot_alea
        liste_mots_deviner.push(mot_alea);
    }
    console.log("Mots générés pour la partie :", liste_mots_deviner);
    return liste_mots_deviner;
}

function jeu(){
    // D'abord demander et paramétrer le nombre de joueurs (int, pas chaine)
    let nb_joueur = parseInt(prompt('Nombre de joueurs ? '));

    let nb_manches = parseInt(prompt('Quel nombre de manche jouer ?'))

    // Ensuite, Charger la liste des mots aléatoires
    let pioche = generation_pioche(nb_manches, 'mots.txt')

    // Ensuite, lancement d'une manche 
    for (let num_manche = 1; num_manche <= nb_manches; num_manche++) {
        manche(nb_joueur, pioche[num_manche - 1]); // Mot de la manche
    }
        // proposition des indices pour chaque joueur
            // vérification de l'indice par rapport au mot secret avec comparaison
            // vérification de l'indice par rapport aux autres indices donnés avec comparaison
            // ajout de l'indice, ou recommencer la démarche de l'indice du joueur
        // lancement du tour pour deviner du joueur
    
        // calcul des scores, et affichage
        // écriture du score dans le fichier
}

function comparaison(mot_a, mot_b){
    // True si mots pareils, ou racines identiques
    if (mot_a === mot_b){
        return true
    }
    //else {
    //    for (let iteration = 0; iteration <mot_a.length; iteration++){
    //        // à voir comment on définit la racine
    //        pass
    //    }
    //} 
    else {
        return false
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
    return liste_indice
}


function score(prompt, mot, pioche, mots_trouves){
    if (prompt===0) {
        console.log("Tu passes");
    }
    else if (verification(mot, prop) === true){
        console.log(`Réussite, c'était bien le mot ${mot}`)
        mots_trouves.push(mot);
    }
    else // Erreur
    {
        if(pioche.length>0){
            console.log(`Echec, le mot était ${mot}.J'enlève une carte de la pioche`);
            pioche.shift(); // Supprime le premier mot de la liste
        }else{ // Plus de mots dans la pioche, on retire une carte des mots trouvés
            console.log(`Echec, le mot était ${mot}.J'enlève une carte des mots trouvés`);
            mots_trouves.shift(); // Supprime le premier mot de la liste
        }
    }
}

function calcul_score(mots_trouves){
    let score = mots_trouves.length
    console.log(`Fin du jeu ! Vous avez ${score} points.`)

    if (score<4){
        console.log("Essayez encore");
    }else if (3<score<7){
        console.log("C'est un bon début. Réessayez !");
    }else if (8<score<11){
        console.log("Vous êtes dans la moyenne. Arriverez-vous à faire mieux ?");
    }else if (score===11){
        console.log("Génial ! C'est un score qui se fête !");
    }else if (score===12){
        console.log("Incroyable ! Vos amis doivent être impressionnés !");
    }else if (score===13){
        console.log("Score parfait ! Y arriverez-vous encore ?");
    }
}

function proposition(indice) {
    console.log("Voici les indices proposés par les joueurs :");
    for(let i;i++;i<indice.length){
        console.log(`Indice ${i} : ${indice[i]}`);
    }
    console.log("Vous avez le droit à un seul essai. Si vous voulez passez, tapez 0");
    let prop = prompt("Votre proposition : ");
    return prop;
}

function ecrireDansFichier(contenu, new_file=true) {
    if (new_file===true) {
        fs.writeFile("historique.txt", contenu, (err)=> {
            if (err) {
                console.log('Erreur lors de l\'écriture dans le fichier :', err);
            } else {
                console.log(`Le contenu a été écrit dans ${nomFichier}`);
            }
        })
    }
}


jeu();