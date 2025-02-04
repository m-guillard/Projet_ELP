# projet_elp
# elm
C:\Users\alice\Documents\Alice\INSA\3A\S1\Projets_2\ELP\Projet_ELP\ELM>elm make src/main.elm --output=index.html

PS C:\Users\alice\Documents\Alice\INSA\3A\S1\Projets_2\ELM_test> elm make src/Main.elm

-> Partie View
    - Partie HTML et CSS réalisée
    - J'ai mis des imports en commentaires pour éviter les erreurs lors de la compilation
    - Génère un fichier index.html qui correspond exactement au site projet
    - Commande pour compiler : elm make Main.elm --output=index.html
    - [ display model ] en commentaires pour pouvoir compiler et tester le fichier, display fait référence au module Display qui affiche le dessin
-> Update : gestion quand on clique sur bouton (envoi vers Parsing)
    - Action à réaliser lorsqu'on clique sur Draw (change le model)
    - Il faudra donc revoir la partie Model
-> Module Parsing (partie logique) : traiter les informations tapées sur la barre

** A voir plus tard
-> Init et update (à voir plus tard)
-> Affichage du dessin à partir de la structure de données de Parsing (à voir plus tard)