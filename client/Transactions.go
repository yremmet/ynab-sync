package client

/*
{
  "data": {
    "transactions": [
      {
        "id": "string",
        "date": "string",
        "amount": 0,
        "memo": "string",
        "cleared": "cleared",
        "approved": true,
        "flag_color": "red",
        "account_id": "string",
        "payee_id": "string",
        "category_id": "string",
        "transfer_account_id": "string",
        "transfer_transaction_id": "string",
        "matched_transaction_id": "string",
        "import_id": "string",
        "deleted": true,
        "account_name": "string",
        "payee_name": "string",
        "category_name": "string",
        "subtransactions": [
          {
            "id": "string",
            "transaction_id": "string",
            "amount": 0,
            "memo": "string",
            "payee_id": "string",
            "payee_name": "string",
            "category_id": "string",
            "category_name": "string",
            "transfer_account_id": "string",
            "transfer_transaction_id": "string",
            "deleted": true
          }
        ]
      }
    ],
    "server_knowledge": 0
  }
}
 */
type Cleared string
type FlagColor string
const (
	TRANSACTION_CLEARED="cleared"
	TRANSACTION_CLEARED_UNCLEARED="uncleared"

	TRANSACTION_FLAG_COLOR_RED    = "red"
	TRANSACTION_FLAG_COLOR_BLUE   = "blue"
	TRANSACTION_FLAG_COLOR_ORANGE = "orange"
	TRANSACTION_FLAG_COLOR_YELLOW = "yellow"
	TRANSACTION_FLAG_COLOR_GREEN  = "green"
	TRANSACTION_FLAG_COLOR_PURPLE = "purple"
)

type Transaction struct {
	Id                    string           `json:"id,omitempty""`
	Amount                int              `json:"amount"`
	Date                  string           `json:"date"`
	Memo                  string           `json:"memo,omitempty"`
	Cleared               Cleared          `json:"cleared,omitempty"`
	Approved              bool             `json:"approved,omitempty"`
	FlagColor             FlagColor        `json:"flag_color,omitempty"`
	AccountId             string           `json:"account_id,omitempty"`
	AccountName           string           `json:"account_name,omitempty"`
	TransactionId         string           `json:"transaction_id,omitempty"`
	PayeeId               string           `json:"payee_id,omitempty"`
	PayeeName             string           `json:"payee_name,omitempty"`
	CategoryId            string           `json:"category_id,omitempty"`
	CategoryName          string           `json:"category_name,omitempty"`
	TransferAccountId     string           `json:"transfer_account_id,omitempty"`
	TransferTranscationId string           `json:"transfer_transaction_id,omitempty"`
	MatchedTransactionId  string           `json:"matched_transaction_id,omitempty"`
	ImportId              string           `json:"import_id,omitempty"`
	Deleted               string           `json:"deleted,omitempty"`
	SubTransactions       []SubTransaction `json:"subtransactions,omitempty"`
}



type SubTransaction struct {
	Id string `json:"id,omitempty""`
	TransactionId string `json:"transaction_id,omitempty"`
	Amount int `json:"amount"`
	Memo string `json:"memo,omitempty"`
	PayeeId string `json:"payee_id,omitempty"`
	PayeeName string `json:"payee_name,omitempty"`
	CategoryId string `json:"category_id,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
	TransferAccountId string `json:"transfer_account_id,omitempty"`
	TransferTranscationId string `json:"transfer_transaction_id,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
}
type Transactions struct {
	Data struct {
		Transactions []Transaction `json:"transactions"`
		TransactionIds []string    `json:"transaction_ids"`
		ServerKnowledge int        `json:"server_knowledge,omitempty"`
	}`json:"data"`
}

type TransactionsEnvelope struct {
	Transactions []Transaction `json:"transactions"`
}

