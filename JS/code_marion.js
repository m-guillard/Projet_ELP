const fs = require('fs');

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
        console.log(`Echec, le mot était ${mot}. J'enlève une carte de la pioche`);
        pioche.shift(); // Supprime le premier mot de la liste
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
        fs.writeFile("historique")
    }
}