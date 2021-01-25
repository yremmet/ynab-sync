package client

/*
{
  "data": {
    "accounts": [
      {
        "id": "415a0e8b-9bb2-40a2-9005-501bd01d5a84",
        "name": "test account",
        "type": "checking",
        "on_budget": true,
        "closed": false,
        "note": null,
        "balance": 47030,
        "cleared_balance": 0,
        "uncleared_balance": 47030,
        "transfer_payee_id": "bcb4fb1f-1ef2-4505-bda0-b2c1cd261a47",
        "deleted": false
      },
      {
        "id": "b59d7b51-088d-4bb5-9cf6-a81ac8e0402f",
        "name": "test account",
        "type": "checking",
        "on_budget": true,
        "closed": false,
        "note": null,
        "balance": 0,
        "cleared_balance": 0,
        "uncleared_balance": 0,
        "transfer_payee_id": "313918dc-06f9-4305-99d9-c1d9260208d4",
        "deleted": false
      }
    ],
    "server_knowledge": 63
  }
}
 */

type Account struct {
	Id string `json:"id,omitempty""`
	Name string `json:"name""`
	Type string `json:"type"`
	OnBudget bool `json:"on_budget,omitempty"`
	Closed string `json:"closed,omitempty"`
	Note string `json:"note,omitempty"`
	Balance int `json:"balance,omitempty"`
	ClearedBalance int  `json:"cleared_balance,omitempty"`
	UnclearedBalance int `json:"uncleared_balance,omitempty"`
	TransferPayeeId string `json:"transfer_payee_id,omitempty"`
}

type Accounts struct {
	Data struct {
		Accounts []Account  `json:"accounts"`
		ServerKnowledge int `json:"server_knowledge,omitempty"`
	}`json:"data"`
}

type AccountEnvelope struct {
	Account Account `json:"account"`
}

