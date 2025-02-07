// PS C:\Users\alice\Documents\Alice\INSA\3A\S1\Projets_2\ELP\Projet_ELP\JS> node code_Alice.js

const prompt = require('prompt-sync')();  // Charger le module prompt-sync
const fs = require('fs-extra');  // Correct
const stringSimilarity = require("string-similarity");
const randomWordFR = require('random-word-fr');


function manche(indics, mot_a_deviner){
    console.log(`\n Début de la maanche.`);

    let [liste_indices, liste_indices_conserves] = indices(indics, mot_a_deviner);
    console.log("\n🔍 Voici la liste des indices proposés :");
    console.log(liste_indices);

    console.log("\n✅ Voici la liste des indices conservés : ");
    console.log(liste_indices_conserves)
    prompt('Appuyer sur Entrée pour lancer la manche.')

    cacher_mots();
    
    return [liste_indices, liste_indices_conserves];

}

function cacher_mots(){
    for (i =0; i<35; i++){
        console.log("\n");
    }
}


function generation_pioche(nombre_mots, fichier){
    let mots = fs.readFileSync(fichier, 'utf8').split('\n').map(mot => mot.trim()); // Correction de FileSystem -> fs
    const liste_mots_deviner = [];
    for (let i = 0; i < nombre_mots; i++){
        // let mot_alea = mots[Math.floor(Math.random() * mots.length)]; // Ajout de mot_alea
        let mot_alea = randomWordFR()
        liste_mots_deviner.push(mot_alea);
    }
    return liste_mots_deviner;
}

function jeu(){
    ecrireDansFichier(`Début jeu \n`, true);
    // D'abord demander et paramétrer le nombre de joueurs (int, pas chaine)
    let nb_joueur = parseInt(prompt('Nombre de joueurs ? '));
    while (nb_joueur <2 || isNaN(nb_joueur)){
        nb_joueur = parseInt(prompt('Il faut au moins 2 joueurs. A combien voulez-vous jouer ? '));
    }

    let noms_joueurs = [];
    for (let i = 0; i < nb_joueur; i++) {
        let nom = prompt(`Nom du joueur ${i + 1} : `);
        noms_joueurs.push(nom);
    }

    let nb_manches = parseInt(prompt('Combien de manches jouer ? '))
    while (isNaN(nb_joueur) || nb_manches <1){
        nb_manches = parseInt(prompt('Il faut au moins jouer une manche. Combien de manches jouer ? '));
    }
    
    console.log("\nLors de ce jeu, nous appelerons 'Devineur' le joueur qui devinera les mots, et les 'Indics' les joueurs chargés de proposer les indices au Devineur.")
    console.log("\n🎲 Chaque joueur sera Devineur une fois avant que la boucle recommence.");
    
    let pioche = generation_pioche(nb_manches, 'mots.txt');
    let mots_trouves = [];
    let anciens_devineurs = [];
    let num_manche = 1;

    // Ensuite, lancement d'une manche 
    while (pioche.length >0){
        console.log(`\n--- Manche ${num_manche} ---`);

        // Si tous les joueurs sont passés, on reset la liste
        if (anciens_devineurs.length === noms_joueurs.length) {
            anciens_devineurs = [];
        }

        // Choisir un Devineur parmi ceux qui n'ont pas encore joué ce cycle
        let candidats = noms_joueurs.filter(nom => !anciens_devineurs.includes(nom));
        let devineur = candidats[Math.floor(Math.random() * candidats.length)];
        anciens_devineurs.push(devineur);

        // Autres joueurs = Indics
        let indics = noms_joueurs.filter(nom => nom !== devineur);

        console.log(`\n🎯 Le Devineur est : ${devineur}`);
        console.log(`💡 Les Indics sont : ${indics.join(', ')}`);

        console.log("\n🎯 Devineur, éloigne-toi !");
        prompt("📢 Indics, appuyez sur Entrée quand vous êtes prêts...");
        console.log(`Le mot à faire deviner pour cette manche est ${pioche[0]}`);

        let [liste_indices, liste_indices_conserves] = manche(indics, pioche[0]); // Mot de la manche
        let prop = proposition(devineur, liste_indices_conserves);
        [pioche, mots_trouves]=score(prop, pioche[0], pioche, mots_trouves, devineur);
        num_manche += 1;
    }
    calcul_score(mots_trouves)
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
function indices(indics, mot_a_deviner) {
    let liste_indice = [];  // Liste où on mettra les indices à chaque manche
    let liste_indices_conserves = [];

    for (let joueur of indics) {
        let indice = prompt(`Joueur ${joueur}, entre ton indice : `);

        // Vérifier que l'indice n'est pas trop similaire au mot à deviner
        while (stringSimilarity.compareTwoStrings(mot_a_deviner, indice)>0.9) {
            console.log(`⚠️ ${joueur}, entre un autre indice, celui-ci est trop similaire au mot à deviner : `);
            indice = prompt(`${joueur}, entre un autre indice : `);
        }
        ecrireDansFichier(`Le joueur ${joueur} donne l'indice ${indice}\n`);
        liste_indice.push(indice);
        cacher_mots();
    }



    for (let i= 0; i < liste_indice.length; i++){
        let compteur = 1;
        for (let j= 0; j < liste_indice.length; j++){
            if (i != j){
                if (liste_indice[i] === liste_indice[j]){
                    compteur += 1;
                }
            }
        }
        if (compteur == 1){
            liste_indices_conserves.push(liste_indice[i]);
        }
    }
    // Afficher tous les indices collectés
    // console.log("\nListe des indices des joueurs :");
    // console.log(liste_indice);
    return [liste_indice, liste_indices_conserves]
}

function score(prop, mot, pioche, mots_trouves, devineur){
    pioche.shift();
    if (prop==="0") {
        console.log("⏭️ Passage du tour.");
    }
    else if (comparaison(mot, prop) === true){
        console.log(`🎉 Bravo ${devineur} ! Le mot était bien "${mot}" !`);
        mots_trouves.push(mot);
    }
    else // Erreur
    {
        console.log(`❌ Mauvaise réponse. Le mot était "${mot}".`);

        if(pioche.length>0){
            console.log(`❌ Echec, le mot était ${mot}.J'enlève une carte de la pioche`);
            pioche.shift(); // Supprime le premier mot de la liste
        }else{ // Plus de mots dans la pioche, on retire une carte des mots trouvés
            console.log(`🔻 Une carte des mots trouvés a été enlevée.`);
            mots_trouves.shift(); // Supprime le premier mot de la liste
        }
    }
    return [pioche, mots_trouves]
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
        console.log("🎊 Génial ! C'est un score qui se fête !");
    }else if (score===12){
        console.log("🔥 Incroyable ! Vos amis doivent être impressionnés !");
    }else if (score===13){
        console.log("🏅 Score parfait ! Y arriverez-vous encore ?");
    }
}

function proposition(devineur, liste_indices_conserves) {
    console.log("\n📢 Voici les indices proposés par les joueurs :");
    if (liste_indices_conserves.length > 0){
        for (let i = 0; i < liste_indices_conserves.length; i++) {
            console.log(`💡 Indice ${i + 1} : ${liste_indices_conserves[i]}`);
        }
    } else {
        console.log("Dommage, pas d'indices pour cette fois (eh oui fallait les varier oupsi)");
    }

    console.log("\n🤔 Vous avez le droit à un seul essai. Si vous voulez passez, tapez 0");
    let prop = prompt(`${devineur}, quelle est ta proposition ? `);
    return prop;
}

async function ecrireDansFichier(contenu, new_file=false) {
    let mode = '';
    if (new_file===true) {
        mode = 'w';
    }else {
        mode = 'a';
    }
    try{
        await fs.writeFile("historique.txt", contenu, {flag: mode});
    } catch(err) {
        console.log('❌ Erreur lors de l\'écriture dans le fichier :', err);
    }

}


jeu();