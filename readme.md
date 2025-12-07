# CRM en Go (Gestionnaire de Contacts)

Ce projet est une application **CLI (ligne de commande)** pour la gestion de contacts, développée en **Go (Golang)**.
Elle permet d’ajouter, lister, mettre à jour et supprimer des contacts, avec persistance des données dans un fichier JSON.

**Membre du Groupe :** VEGBA LUX

## 🌳 Historique des Branches

| Version | Branche | Description |
| :--- | :--- | :--- |
| **V1** | `main` | Création initiale du CRM |
| **V2** | `Refactoring_v2_du_CRM` | Refactoring et amélioration de l’architecture |
| **V2.2** | `Refactoring_v2.2_du_CRM` | Passage à une architecture services / handlers / repository |
| **V3** | `Main` | Intégration complète avec Cobra et persistance JSON |

---

## ▶️ Prérequis

* **Go 1.25** ou supérieur installé
* **Git** (pour cloner le dépôt)
* **Terminal** : PowerShell (Windows) ou Terminal (Linux/macOS)

---

## ▶️ Installation et Lancement

1.  **Cloner le dépôt :**

    ```bash
    git clone [https://github.com/iimAtomic/CRM-en-GO.git](https://github.com/iimAtomic/CRM-en-GO.git)
    cd CRM-en-GO
    ```

2.  **Installer les dépendances :**

    ```bash
    go mod tidy
    ```

3.  **Lancer l’application en CLI :**

    ```bash
    go run main.go --help
    ```

    Vous verrez toutes les commandes disponibles s'afficher :

    ```text
    Usage:
      crm [command]

    Available Commands:
      add          Ajouter un contact
      delete       Supprimer un contact
      list         Lister tous les contacts
      update       Mettre à jour un contact
      help         Help about any command
    ```

---

## ▶️ Commandes principales

| Commande | Description |
| :--- | :--- |
| `crm add` | Ajouter un nouveau contact |
| `crm list` | Lister tous les contacts |
| `crm update` | Mettre à jour un contact existant |
| `crm delete` | Supprimer un contact par ID |

> 📁 **Note :** Les contacts sont enregistrés dans le fichier `contacts.json` à la racine du projet.

---

## ▶️ Générer un exécutable

Vous pouvez compiler le projet pour l'utiliser sans la commande `go run`.

### Pour Windows

Créer le fichier `.exe` :
```powershell
go build -o crm.exe main.go
```
PowerShell

```
.\crm.exe --help
```
Pour Linux / macOS
Créer l'exécutable :

```Bash

go build -o crm main.go

```
Lancer l'application :

```Bash

./crm --help


Architecture du projet :

```Text
Le projet suit une architecture structurée pour séparer les responsabilités :
```

- cmd/ : Commandes Cobra (add, list, update, delete).
- handlers/ : Gestion des interactions avec l’utilisateur (entrées/sorties).
- service/ : Logique métier de l'application.
- repository/json/ : Gestion de la persistance des contacts (lecture/écriture JSON).
- models/ : Définition des structures de données (ex: Modèle Contact).