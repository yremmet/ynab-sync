package client

import (
	"encoding/json"
	"fmt"
)

type TransactionsService struct {
	client *YNABClient
}

func NewTransactionsService (client *YNABClient) TransactionsService {
	return TransactionsService{client: client}
}

func (service *TransactionsService) GetTransactions(budget Budget) Transactions {
	body := service.client.sendRequest(fmt.Sprintf("/budgets/%s/transactions", budget.Id), "GET", nil)
	var acc Transactions
	json.Unmarshal(body, &acc)
	return acc
}

func (service *TransactionsService) CreateTransactions(budget Budget, transactions *[]Transaction) {
	t := TransactionsEnvelope{Transactions: *transactions}
	x,_ := json.Marshal(t)
	service.client.sendRequest(fmt.Sprintf("/budgets/%s/transactions", budget.Id) , "POST", x)
}
