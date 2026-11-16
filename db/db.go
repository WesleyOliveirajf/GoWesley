package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConectaBancodeDados() *sql.DB {
	senha := os.Getenv("DB_PASSWORD")
	conexao := fmt.Sprintf("user=wesleyoliveira dbname=Alura_Loja password=%s host=localhost port=5433 sslmode=disable connect_timeout=5", senha)
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
