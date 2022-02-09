package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Drinnn/kool-bank/infrastructure/repository"
	"github.com/Drinnn/kool-bank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()
}

func setupTransactionUseCase(db *sql.DB) *usecase.TransactionUseCase {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewTransactionUseCase(transactionRepository)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error trying to connect to database")
	}

	return db
}
