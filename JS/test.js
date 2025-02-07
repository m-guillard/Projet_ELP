const prompt = require('prompt-sync')();  
const fs = require('fs');  

function jeu(){
    let nb_joueur = parseInt(prompt('Nombre de joueurs ? '));
    while (nb_joueur < 2 || isNaN(nb_joueur)) {
        nb_joueur = parseInt(prompt('Il faut au moins 2 joueurs. A combien voulez-vous jouer ? '));
    }

    let noms_joueurs = [];
    for (let i = 0; i < nb_joueur; i++) {
        let nom = prompt(`Nom du joueur ${i + 1} : `);
        noms_joueurs.push(nom);
    }

    let nb_manches = parseInt(prompt('Combien de manches jouer ? '));
    while (isNaN(nb_manches) || nb_manches < 1) {
        nb_manches = parseInt(prompt('Il faut au moins jouer une manche. Combien de manches jouer ? '));
    }

    console.log("\n🎲 Chaque joueur sera Devineur une fois avant que la boucle recommence.");
    
    let pioche = generation_pioche(nb_manches, 'mots.txt');
    let mots_trouves = [];
    let anciens_devineurs = [];

    for (let num_manche = 1; num_manche <= nb_manches; num_manche++) {
        console.log(`\n--- Manche ${num_manche} ---`);

        // Si tous les joueurs sont passés, on reset la liste
        if (anciens_devineurs.length === noms_joueurs.length) {
            anciens_devineurs = [];
        }

        // Choisir un Devineur parmi ceux qui n'ont pas encore joué ce cycle
        let candidats = noms_joueurs.filter(nom => !anciens_devineurs.includes(nom));
        let devineur = candidats[Math.floor(Math.random() * candidats.length)];
        anciens_devineurs.push(devineur);

        // Les autres joueurs sont les Indics
        let indics = noms_joueurs.filter(nom => nom !== devineur);

        console.log(`\n🎯 Le Devineur est : ${devineur}`);
        console.log(`💡 Les Indics sont : ${indics.join(', ')}`);

        let liste_indices = manche(indics, pioche[num_manche - 1]);
        let prop = proposition(devineur, liste_indices);
        score(prop, pioche[num_manche - 1], pioche, mots_trouves, devineur);
    }

    calcul_score(mots_trouves);
}

function manche(indics, mot_a_deviner){
    console.log("\nDébut de la manche.");
    let validation = 'n';
    let liste_indices = [];

    while (validation !== 'y'){
        liste_indices = indices(indics, mot_a_deviner);

        console.log("\n🎯 Devineur, éloigne-toi !");
        prompt("📢 Indics, appuyez sur Entrée quand vous êtes prêts...");
        console.log("\n🔍 Voici la liste des indices proposés :");
        console.log(liste_indices);

        validation = prompt("✅ Les validez-vous [y/n] ? ");
        if (validation !== 'y') {
            console.log("❌ Recommençons la phase des indices...");
        }
        cacher_mots();
    }

    return liste_indices;
}

function indices(indics, mot_a_deviner) {
    let liste_indice = [];

    for (let joueur of indics) {
        let indice = prompt(`${joueur}, entre ton indice : `);

        while (comparaison(mot_a_deviner, indice)) {
            console.log(`⚠️ ${joueur}, ton indice est trop similaire au mot secret.`);
            indice = prompt(`${joueur}, entre un autre indice : `);
        }

        let estUnique = true;
        for (let i = 0; i < liste_indice.length; i++) {
            if (comparaison(liste_indice[i], indice)) {
                estUnique = false;
                console.log(`⛔ ${joueur}, cet indice a déjà été donné.`);
                break;
            }
        }

        if (estUnique) {
            liste_indice.push(indice);
        } else {
            console.log(`🔁 ${joueur}, entre un nouvel indice.`);
        }
    }

    return liste_indice;
}

function proposition(devineur, indices) {
    console.log("\n📢 Voici les indices donnés :");
    indices.forEach((indice, i) => {
        console.log(`💡 Indice ${i + 1} : ${indice}`);
    });

    console.log("\n🤔 Tu as un seul essai. Tape 0 pour passer.");
    let prop = prompt(`${devineur}, quelle est ta proposition ? `);
    return prop;
}

function comparaison(mot_a, mot_b){
    return mot_a === mot_b; // Simple pour l'instant, peut être amélioré
}

function generation_pioche(nombre_mots, fichier){
    let mots = fs.readFileSync(fichier, 'utf8').split('\n').map(mot => mot.trim());
    return Array.from({length: nombre_mots}, () => mots[Math.floor(Math.random() * mots.length)]);
}

function cacher_mots(){
    console.log("\n".repeat(35)); // Nettoie l'écran
}

function score(prop, mot, pioche, mots_trouves, devineur){
    if (prop === "0") {
        console.log("⏭️ Passage du tour.");
        ecrireDansFichier(`${devineur} passe.`);
    } else if (comparaison(mot, prop)) {
        console.log(`🎉 Bravo ${devineur} ! Le mot était bien "${mot}" !`);
        mots_trouves.push(mot);
        ecrireDansFichier(`${devineur} a trouvé le mot "${mot}".`);
    } else {
        console.log(`❌ Mauvaise réponse. Le mot était "${mot}".`);
        ecrireDansFichier(`${devineur} a proposé "${prop}", mais le mot était "${mot}".`);

        if (pioche.length > 0) {
            console.log("🔻 Suppression d'une carte de la pioche.");
            pioche.shift();
        } else {
            console.log("🔻 Suppression d'un mot trouvé.");
            mots_trouves.shift();
        }
    }
}

function calcul_score(mots_trouves){
    let score = mots_trouves.length;
    console.log(`\n🏆 Fin du jeu ! Score final : ${score} points.`);
    ecrireDansFichier(`Fin du jeu ! Score : ${score} points.`);

    if (score < 4) {
        console.log("Essayez encore !");
    } else if (score < 7) {
        console.log("Bon début !");
    } else if (score < 11) {
        console.log("Vous êtes dans la moyenne !");
    } else if (score === 11) {
        console.log("🎊 Score à fêter !");
    } else if (score === 12) {
        console.log("🔥 Impressionnant !");
    } else if (score === 13) {
        console.log("🏅 Score parfait !");
    }
}

function ecrireDansFichier(contenu, new_file=false) {
    contenu += "\n";
    if (new_file) {
        fs.writeFile("historique.txt", contenu, (err) => {
            if (err) console.log('❌ Erreur :', err);
        });
    } else {
        fs.appendFile("historique.txt", contenu, (err) => {
            if (err) console.log('❌ Erreur :', err);
        });
    }
}

jeu();
