package client

import (
	"bytes"
	"fmt"
	"github.com/yremmet/ynab/config"
	"io/ioutil"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type YNABClient struct {
	bearerToken         string
	baseURL             string
	BudgetsService      BudgetsService
	AccountsService     AccountsService
	TransactionsService TransactionsService
	httpClient          HTTPClient
}
func (ynab *YNABClient) SetClient(client HTTPClient){
	ynab.httpClient = client
}


func NewYNABClient(config config.Config) YNABClient {
	client := YNABClient{bearerToken: config.PersonalAccessToken,
		baseURL: "https://api.youneedabudget.com/v1/",
	}
	client.BudgetsService  = NewBudgetsService(&client)
	client.AccountsService = NewAccountsService(&client)
	client.TransactionsService = NewTransactionsService(&client)

	client.httpClient = &http.Client{
		Timeout: config.HttpTimeout,
	}
	return client;
}

func (ynab *YNABClient) sendRequest(path string, method string, body []byte) ([]byte, *ClientError) {
	client := &http.Client{}
	// Create request
	b := bytes.NewReader(body)

	req, _ := http.NewRequest(method, ynab.baseURL+path, b)
	req.Header.Add("Authorization", "bearer "+ ynab.bearerToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failure : ", err)
		return nil, defaultError(resp, err)
	}

	if (resp.StatusCode > 201) {
		return nil, httpError(resp)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody, nil
}