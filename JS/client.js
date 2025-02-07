const net = require('net');
const prompt = require('prompt-sync')();

// Step 1 : Récupérer port et serveur où veut se connecter le joueur
// Step 2 : Connexion au serveur
// Step 3 : Reception début jeu
// BOUCLE
// Step 3.9 : Reception si indic ou devineur
// Step 4 : Si joueur doit deviner
// Step 5 : Recoit message qui lui dit d'attendre les autres
// Step 6 : Recoit les mots
// Step 7 : Envoi d'une proposition
// Step 8 : Affichage de si il a réussi ou non

// Step 4 : Si joueur doit faire deviner
// Step 5 : Envoie mot
// Step 6: Message d'attendre que les indices soient envoyés et que le devineur devine
// Step 7 : Affichage si le joueur a réussi ou non

// Step finale : AFFICHAGE SCORE GLOBAL
function jeu()
{
    
}


function connexion_serveur() {
    let num_port = Nan;
    while (num_port === NaN){
        let serveur = prompt("Nom du serveur : ");
        num_port = parseInt(prompt("Numéro de port : "));
    }
    const client = net.createConnection({ port: port, host: serveur }, () => {
        console.log('Connected to server!');
        jeu();
        client.write('Hello from client!');
    });
    
    client.on('data', (data) => {
        console.log('Received from server:', data.toString());
        client.end(); // Close the connection
    });
    
    client.on('end', () => {
        console.log('Disconnected from server');
    });

}
