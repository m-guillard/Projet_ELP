# Projet Go - Distance de Levensthein

## Description du projet
Le but de ce projet est de comparer des bases de données où les données ont été rentrées à la main. Grâce à la distance de Levenshtein, on peut identifier les erreurs de frappe et ainsi en déduire que deux champs sont identiques.
Ce projet prend en entrée deux bases de données sous un format CSV ainsi que le nom des colonnes à comparer. Il renvoie dans un nouveau fichier CSV, les valeurs de chaque bases de données qui ont une distance de Levenshtein faible (choisie par l'utilisateur).

## Lancer le projet
Pour lancer le serveur : go run tcp_serveur.go Code.go NUMERO_PORT

Pour lancer le client : go run tcp_client.go ADRESSE_SERVEUR NUMERO_PORT NOM_BDD1 NOM_COLONNE_BDD1 NOM_BDD2 NOM_COLONNE_BDD2 NOMBRE_GOROUTINES DISTANCE_LEVENSHTEIN_MAX


## Contributeurs
Alice INVERNIZZI et Marion GUILLARD
