# Documentation

## Pré-requis

### OS

Avant de commencer à développer l'api, il faut tout d'abord mettre en place l'environnement pour le faire.

Nous commençons donc par installer une distribution de Linux tel que Debian pour des raisons pratiques liées au différents composant de l'api mais aussi aux outils qui seront utilisés.
Pour cela, nous passons par WSL afin de pouvoir installer Debian:
```wsl install -d Debian```
Il est tout à fait possible d'installer une autre distribution de Linux, le choix de Debian est purement une préférence.

Après cela, nous installons Docker et GO

### Docker

Pour installer Docker, nous allons suivre la documentation officiel: https://docs.docker.com/engine/install/debian/
Pour des raisons de faciliter, voici quelques étape supplémentaire à effectuer:
```sudo groupadd -f docker
sudo chown root:docker /var/run/docker.sock
sudo usermod -a -G docker "$(whoami)"
newgrp docker
sudo service docker start
```

Passons à présent à Go

### GO

Pour installer GO, nous allons d'abord télécharger l'archive sur le site officiel puis suivre les instructions: https://go.dev/doc/install

Passons maintenant à la partie Kubernetes et Helm.

### Kubernetes

Pour utiliser un cluster Kubernetes, nous avons besoin de kubectl et d'un cluster Kubernetes.
Nous allons utiliser Minikube pour créer notre cluster pour des raisons de simplicité et de contraintes matérielles.

Pour installer Kubectl et Minikube, il suffit de suivre la documentation officiel:
- https://kubernetes.io/fr/docs/tasks/tools/install-minikube/
- https://kubernetes.io/fr/docs/tasks/tools/install-kubectl/#install-kubectl-on-linux
Pour ma part, j'ai utilisé le binaire autonome pour des raisons de facilité.

### Helm

Pour "transformer" les manifests Kubernetes de l'api server en chart Helm, nous allons tout d'abord installer Helm.
Pour ce faire, il suffit de suivre la procédure d'installation du site officiel:
```curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
```

## API

Pour envoyer une requête POST, nous utilisons la commande suivante:
```curl -X POST -d '{"timestamp": "2024-07-01T12:00:00Z", "heartbeat": 75}' -H "Content-Type: application/json" http://localhost:8080/metrics```
Il est possible de changer la valeur qui est dans l'exemple ci-dessus.


Pour obtenir des valeur dans un interval défini avec une requête GET, nous utilisons la commande suivante:
```curl "http://localhost:8080/metrics?start=2024-07-01T00:00:00Z&end=2024-07-30T00:00:00Z"```
L'interval de temps est défini sous la forme suivante: start=AAAA-MM-JJTHH-MM-SSZ&end=AAAA-MM-JJTHH-MM-SSZ.
Les valeur à modifier se situent après ```start``` et ```end```

## Résultats

Lorsque l'on envoie des requêtes POST, le résultat de cette dernière est stocké dans la base de donnée sqlite en JSON.
En voici un exemple:
```json
[
  {"ID":1,"timestamp":"2024-07-01T12:00:00Z","heartbeat":75},
  {"ID":2,"timestamp":"2024-07-01T12:00:00Z","heartbeat":75},
  {"ID":3,"timestamp":"2024-07-02T16:30:00Z","heartbeat":150}
]
```

## Problèmes rencontrés

### BDD

L'utilisation d'une base de données s'est révélé plus compliqué que prévu une fois arrivé au stade du déploiement sur le cluster Kubernetes. La base de données utilisé qui était du sqlite3, était continuellement dans l'état `CrashLoopBackOff`. Malgré plusieurs tentatives, une solution n'a pas été trouvé. J'ai donc décidé de changer pour du MySQL.

### L'api n'est pas `Ready`

L'api semble ne pas passé à l'état `Ready` même si cette dernière semble bien fonctionné et est `Running`. Cela ne semble pas être un soucis pour l'instant mais devra être solutionné plus tard.

## Suite du projet

Par la suite, il serait intéressant d'ajouter la stack monitoring afin de surveiller les metrics de l'api.
Il serait également judicieux de revoir les images docker utilisés afin de réduire le pods des images, les simplifiés et ajouter des mécanismes afin d'assuré le bon déploiement de l'api.
De même, au niveau du cluster, nous pourrions ajouter des sondes, un ingress, définir les ressources et mieux séparer l'api et la base de données en deux charts diférents qui interagisseront entre eux une fois déployé.
