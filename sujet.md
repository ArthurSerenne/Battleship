TP Final – Bataille navale temps réel
Sujet
Vous devez écrire un jeu de bataille navale se jouant en temps réel.
​

Règle du jeu
La bataille navale se joue sur une grille d’au moins 10x10.
​

Le joueur doit deviner où se situent les bateaux adverses afin de les couler.
​

Les bateaux sont placés aléatoirement sur les grilles des joueurs.
​

Une case ne peut être occupée que par un seul morceau de bateau.
​

Il doit y avoir plusieurs bateaux présents sur le plateau.
​

Particularité :

Le jeu se joue en temps réel, ce n’est pas un jeu au tour par tour.
​

N’importe quel joueur peut jouer plusieurs fois d’affilée sans attendre les autres.
​

Programme
Un seul binaire existe : le programme agit à la fois comme serveur et client.
​

Chaque joueur lance l’exécutable sur sa machine pour jouer.
​

S’il y a 6 joueurs, le binaire doit être exécuté 6 fois.
​

Partie serveur
Au lancement du programme, un nouveau plateau est généré : c’est la zone navale personnelle du joueur.
​

Le serveur doit exposer plusieurs routes HTTP :
​

GET /board
Renvoie l’état du plateau.
​

Exemple de format (à adapter si besoin) :
​

0 : case intouchée, non dévoilée

1 : coup dans l’eau

2 : bateau touché

Vous pouvez choisir un meilleur format de communication si nécessaire.
​

GET /boats
Renvoie le nombre de bateaux encore à flot.
​

Si la route renvoie 0, le joueur a perdu.
​

Un bateau est à flot si au moins un de ses morceaux n’a pas été touché.
​

POST /hit
Reçoit la case choisie par l’adversaire.
​

Répond le résultat du coup : dans l’eau ou bateau touché.
​

GET /hits
Renvoie les coordonnées des tirs reçus et leur impact.
​

Partie client
Au lancement, on passe l’adresse d’un ou plusieurs programmes adverses (argument ou entrée standard).
​

Le joueur interagit en ligne de commande pour tirer sur un plateau adverse via les routes exposées par l’ennemi.
​

Si le joueur n’a plus de bateaux, il ne peut plus tirer et est considéré comme perdant.
​

Le terminal doit être :
​

Le plus interactif possible.

Utiliser tout l’espace de la fenêtre pour afficher les informations de la partie.

Contraintes
Un seul et unique binaire agit comme serveur et client de jeu asynchrone.
​

On doit pouvoir passer plus d’une adresse au lancement ou en cours de partie.
​

Le jeu doit permettre de jouer à 2, 3 ou plus de joueurs sans difficulté.
​

Bonus
Réaliser au moins 4 bonus en plus du sujet de base, sans interférer avec les points obligatoires.
​

Exemples de bonus possibles :
​

Bateaux en mouvement permanent ou après chaque tir.

Autres types de tirs (en croix, en cercle, paralysant, etc.).

Définition d’un profil joueur et échanges de messages personnalisés.

Partie en plusieurs manches.

Événements aléatoires sur la carte (tornade, tempête, kraken, etc.).

Modalités
Projet à réaliser et à soutenir par groupe de 2 ou 3.
