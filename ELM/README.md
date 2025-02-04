# Projet ELM - TcTurtle

## Description du projet
Le but de ce projet est de générer une page HTML qui prend en entrée une commande en TcTurtle. Le programme génère un dessin en fonction de l'entrée.

## Lancer le projet
Pour générer HTML :  elm make  src/main.elm --output=index.html
Sur la page HTML, index.html, exemple d'entrée :
- [Repeat 360 [ Right 1, Forward 1]]
- [Forward 100, Repeat 4 [Forward 50, Left 90], Forward 100]
- [Repeat 36 [Right 10, Repeat 8 [Forward 25, Left 45]]]
- [Repeat 8 [Left 45, Repeat 6 [Repeat 90 [Forward 1, Left 2], Left 90]]]

## Contributeurs
Alice INVERNIZZI et Marion GUILLARD