# Ava Tool

Ava Tool est un outil en ligne de commande (CLI) conçu pour accélérer la création de services en Go. Il permet de générer une structure de service prête à l’emploi, incluant un routeur (basé sur Gin), une configuration ORM avec Gorm (en cas d’utilisation d’une base de données PostgreSQL), ainsi qu’une intégration Docker, Docker Compose et une configuration CI GitLab.

## Table des matières
- [Contexte et Objectifs](#contexte-et-objectifs)
- [Fonctionnalités](#fonctionnalités)
- [Installation et Prérequis](#installation-et-prérequis)
- [Utilisation](#utilisation)
- [Architecture et Fonctionnement](#architecture-et-fonctionnement)
- [Support](#support)
- [Historique et Versions](#historique-et-versions)
- [Remerciements et Licence](#remerciements-et-licence)
- [Contact](#contact)

## Contexte et Objectifs

Ava Tool répond au besoin de réduire le temps de mise en place d'un service en Go, une tâche souvent répétitive et chronophage.  
L'outil s'adresse principalement aux développeurs ou à toute personne souhaitant initier rapidement un service avec une configuration et une architecture standardisées.

**Principaux avantages :**
- **Gain de temps:** Automatisation de la création d’un service Go complet.
- **Standardisation:** Génération d’une structure de dossiers préconfigurée avec le support d’un routeur (Gin) et d’une base de données (Gorm avec PostgreSQL en option).
- **Intégration DevOps:** Configuration Docker, Docker Compose et CI GitLab pour faciliter le déploiement et l’intégration continue.
- **Extensibilité future:** Possibilité d’ajouter de nouveaux services et alternatives (p.ex. autres protocoles, outils de CI/CD) au fur et à mesure de l'évolution du projet.

## Fonctionnalités

- **Génération de service Go**  
  Crée automatiquement un service Go avec une structure de fichiers adaptée.

- **Configuration du service**
  - Intégration de Gin pour le routage HTTP/REST.
  - Option d’inclure Gorm et une base de données PostgreSQL.
  - Définition du port de service, avec la possibilité d’utiliser une valeur par défaut (8080).

- **Environnement de déploiement**
  - Génération d’une configuration Docker et Docker Compose.
  - Création d’une CI GitLab prête à l’emploi pour automatiser les builds et tests.

- **Commande interactive**
  - L’outil propose une interface en ligne de commande qui pose trois questions clés :
    1. Le nom du service (avec une valeur par défaut « new service »).
    2. L’option d’inclure une base de données PostgreSQL (y/n).
    3. Le port sur lequel le service doit tourner (par défaut 8080).

## Installation et Prérequis

### Prérequis

- **Exécution:**  
  Ava Tool est un binaire compilé, il ne nécessite pas d’environnement particulier pour fonctionner.
  
- **Développement:**  
  Pour développer ou modifier l’outil, assurez-vous d’avoir installé :
  - Go
  - Docker

### Installation

L’outil sera disponible sous forme de release sur GitHub.  
Téléchargez la version correspondant à votre système et ajoutez-la à votre PATH selon les instructions fournies.

## Utilisation

Ava Tool fonctionne en mode interactif. Pour lancer l’outil, exécutez simplement la commande suivante dans votre terminal :

```bash
$ ava-tool
```

Vous serez invité à répondre aux questions suivantes :

1. **Nom du service**  
   (par défaut "new service")  
   Ex : `my-pretty-service`

2. **Inclusion de PostgreSQL**  
   Voulez-vous ajouter une base de données PostgreSQL ?  
   (Répondez `y` pour oui ou `n` pour non)

3. **Port du service**  
   (par défaut 8080)  
   Ex : `8080`

Une fois ces étapes complétées, l’outil génère l’architecture du service.  
Pour démarrer le service, utilisez :

```bash
$ docker-compose up
```

Vous pouvez ensuite vérifier le bon fonctionnement du service en effectuant une requête à l’endpoint `/health`.

## Architecture et Fonctionnement

Le projet est organisé en deux modules principaux :

- **Module `commands`:**  
  Gère l'interaction avec l'utilisateur et la logique des commandes.

- **Module `generator`:**  
  Contient les templates et la logique de génération de la structure de fichiers et de configuration du service.

## Support

Pour signaler un bug ou poser une question, merci d'ouvrir une issue sur le dépôt GitHub du projet.

## Historique et Versions

- **Version 1.0**  
  Première version d'Ava Tool.  
  Cette version inclut la génération de service Go avec support REST et HTTP, ainsi que les configurations pour Docker, Docker Compose et CI GitLab.

## Remerciements et Licence

**Remerciements:**  
Merci aux bibliothèques et aux projets suivants qui ont inspiré ce travail :
- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [GORM](https://gorm.io)
- [Go](https://golang.org)

**Licence:**  
La licence du projet n'a pas encore été définie. Veuillez consulter ce répertoire pour les mises à jour futures concernant la licence.

## Contact

Pour suivre l’évolution du projet ou pour toute question, vous pouvez me contacter via LinkedIn :  
[Alexyvda](https://www.linkedin.com/in/alexyvda/)
