package main

import (
	"fmt"

	"github.com/gaabrieljesuss/bank_sdk/functions"

	"github.com/labstack/gommon/log"
)

func main() {
	baseURL := "http://localhost:8000"
	account, err := functions.GetAccountById(baseURL, 1)
	functions.Deposit(baseURL, 1, 100)
	functions.Withdraw(baseURL, 1, 2)
	functions.GetAllAccount(baseURL)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(account.Name)

}
