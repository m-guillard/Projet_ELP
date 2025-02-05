const fs = require('fs');

function score(prompt, mot, pioche, mots_trouves, joueur){
    if (prompt===0) {
        console.log("Tu passes");
        ecrireDansFichier(`Le joueur ${joueur} passe.`);
    }
    else if (verification(mot, prop) === true){
        console.log(`Réussite, c'était bien le mot ${mot}`);
        mots_trouves.push(mot);
        ecrireDansFichier(`Le joueur ${joueur} a trouvé le mot ${mot}.`);
    }
    else // Erreur
    {
        ecrireDansFichier(`Le joueur ${joueur} a proposé le mot ${prompt}.`);
        ecrireDansFichier(`Le mot à trouver était ${mot}.`);
        if(pioche.length>0){
            console.log(`Echec, le mot était ${mot}.J'enlève une carte de la pioche`);
            pioche.shift(); // Supprime le premier mot de la liste
            ecrireDansFichier(`Une carte de la pioche a été enlevée.`)
        }else{ // Plus de mots dans la pioche, on retire une carte des mots trouvés
            console.log(`Une carte des mots trouvés a été enlevée.`);
            mots_trouves.shift(); // Supprime le premier mot de la liste
        }
    }
}

function calcul_score(mots_trouves){
    let score = mots_trouves.length
    console.log(`Fin du jeu ! Vous avez ${score} points.`)
    ecrireDansFichier(`Fin du jeu ! Les joueurs ont un total de ${score} points.`)

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

function ecrireDansFichier(contenu, new_file=false) {
    contenu += "\n";
    if (new_file===true) {
        fs.writeFile("historique.txt", contenu, (err)=> {
            if (err) {
                console.log('Erreur lors de l\'écriture dans le fichier :', err);
            } else {
                console.log(`Le contenu a été écrit dans ${nomFichier}`);
            }
        })
    }else {
        fs.appendFile("historique.txt", contenu + '\n', (err) => {
            if (err) {
                console.log('Erreur lors de l\'ajout dans le fichier :', err);
            } else {
                console.log(`Le texte a été ajouté dans historique.txt`);
            }
        });
    }

}