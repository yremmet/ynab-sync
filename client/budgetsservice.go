package client

import (
	"encoding/json"
)

type BudgetsService struct {
	client *YNABClient
}

func NewBudgetsService (client *YNABClient) BudgetsService {
	return BudgetsService{client: client}
}

func (service *BudgetsService) GetBudgets() Budgets {
	body := service.client.sendRequest("budgets?include_accounts", "GET", nil)
	var budgets Budgets
	json.Unmarshal(body, &budgets)
	return budgets
}


