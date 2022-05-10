package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gabosarmiento/gogo/adapter/repository"
	"github.com/gabosarmiento/gogo/usecase/execute_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello, World!")

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := execute_transaction.NewExecuteTransaction(repo)

	input := execute_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    2000,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	outputJson, _ := json.Marshal(output)

	fmt.Println(string(outputJson))

}
