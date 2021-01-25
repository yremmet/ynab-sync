package client

import (
	"encoding/json"
	"fmt"
)

type AccountsService struct {
	client *YNABClient
}

func NewAccountsService (client *YNABClient) AccountsService {
	return AccountsService{client: client}
}

func (service *AccountsService) GetAccounts(budget Budget) Accounts {
	body := service.client.sendRequest(fmt.Sprintf("/budgets/%s/accounts?last_knowledge_of_server", budget.Id), "GET", nil)
	var acc Accounts
	json.Unmarshal(body, &acc)
	return acc
}

func (service *AccountsService) CreateAccount(budget Budget, name string, accounttype string, balance int, iban string) {
	a := Account{
		Name:             name,
		Type:             accounttype,
		Note:             iban,
		Balance:          balance,
	}
	ae := AccountEnvelope{Account: a}
	x,_ := json.Marshal(ae)
	service.client.sendRequest(fmt.Sprintf("/budgets/%s/accounts?last_knowledge_of_server",budget.Id) , "POST", x)
}
