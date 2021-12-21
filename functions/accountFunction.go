package functions

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gaabrieljesuss/bank_sdk/domain"

	"github.com/labstack/gommon/log"
)

func GetAccountById(baseURL string, id int) (*domain.Account, *error) {

	url := fmt.Sprintf(baseURL+"/api/account/%d/find", id)

	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return nil, &err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, &err
	}

	account := new(domain.Account)
	json2.Unmarshal(body, &account)

	return account, nil
}

func Deposit(baseURL string, id int, balance float64) *error {
	url := fmt.Sprintf(baseURL+"/api/account/%d/deposit", id)

	json, err := json2.Marshal(map[string]float64{
		"balance": balance,
	})
	if err != nil {
		log.Error(err)
		return &err
	}

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))
	if err != nil {
		log.Error(err)
		return &err
	}
	request.Header.Set("content-type", "application/json")

	client := new(http.Client)
	client.Do(request)

	return nil
}

func Withdraw(baseURL string, id int, balance float64) *error {
	url := fmt.Sprintf(baseURL+"/api/account/%d/withdraw", id)

	json, _ := json2.Marshal(map[string]float64{
		"balance": balance,
	})

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))
	if err != nil {
		log.Error(err)
		return &err
	}
	request.Header.Set("content-type", "application/json")

	client := new(http.Client)
	client.Do(request)

	return nil
}

func GetAllAccount(baseURL string) ([]domain.Account, *error) {

	url := fmt.Sprintf(baseURL + "/api/account/fetch")

	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return nil, &err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, &err
	}

	accounts := []domain.Account{}
	json2.Unmarshal(body, &accounts)

	return accounts, nil
}
