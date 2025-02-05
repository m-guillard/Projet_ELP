const net = require('net');

const client = net.createConnection({ port: 8080, host: '127.0.0.1' }, () => {
    console.log('Connected to server!');
    client.write('Hello from client!');
});

client.on('data', (data) => {
    console.log('Received from server:', data.toString());
    client.end(); // Close the connection
});

client.on('end', () => {
    console.log('Disconnected from server');
});