const net = require('net');

// Step 1 : Récupérer nombre de joueurs et de cartes dans la pioche
// Step 2 : Lancer le serveur et attente que tous les clients rejoignent
// Step 3 : Envoi msg début du jeu
// BOUCLE
// Step 4 : Définir devineur et indics
// Step 5 : Envoi mot au indics / attente au devineur
// Step 6 : Réception indices et vérification validité
// Step 7 : Envoi indices au devineur
// Step 8 : Réception de sa proposition et vérification
// Step 9 : Envoi du résultats aux joueurs

// Step finale : ENVOI SCORE GLOBAL

const server = net.createServer((socket) => {
    console.log('Client connected');

    socket.on('data', (data) => {
        console.log('Received:', data.toString());
        socket.write('Hello from server!');
    });

    socket.on('end', () => {
        console.log('Client disconnected');
    });
});

server.listen(8080, '127.0.0.1', () => {
    console.log('Server listening on port 8080');
});