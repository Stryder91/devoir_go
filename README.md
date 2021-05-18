# Devoir GO

Lionel Tran & Xavier Iyeze

## ORM 
- Nous n'avons pas utilisé ent ORM de Facebook car nous n'avons pas besoin d'une base de données orientée graphe.

Nous avons utilisé sqlx de jmoiron qui s'appuie sur database/sql. Doc : https://jmoiron.github.io/sqlx/

Installation :
$ go get github.com/jmoiron/sqlx
$ go get github.com/mattn/go-sqlite3

## Migration
Pour faire les migrations nous avons utilisé goose dans un premier temps et crée nos schema selon
les énoncés du README.md -> Au final c'était plus facile avec golang-migration (erreur avec goose, 
mais le fichier est toujours dans migrations/00002_v1_v2.go)



