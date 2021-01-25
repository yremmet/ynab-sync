package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type YNABClient struct {
	bearerToken         string
	baseURL             string
	BudgetsService      BudgetsService
	AccountsService     AccountsService
	TransactionsService TransactionsService
}

func NewYNABClient(bearerToken string) (YNABClient){
	client := YNABClient{bearerToken: bearerToken,
		baseURL: "https://api.youneedabudget.com/v1/",
	}
	client.BudgetsService  = NewBudgetsService(&client)
	client.AccountsService = NewAccountsService(&client)
	client.TransactionsService = NewTransactionsService(&client)

	return client;
}

func (ynab *YNABClient) sendRequest(path string, method string, body []byte) []byte {
	client := &http.Client{}
	// Create request
	b := bytes.NewReader(body)

	req, _ := http.NewRequest(method, ynab.baseURL+path, b)
	req.Header.Add("Authorization", "bearer "+ ynab.bearerToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failure : ", err)
	}
	if (resp.Status == "400" ) {
		fmt.Println(string(body))
		fmt.Println(resp)
	}


	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody
}