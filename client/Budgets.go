package client

type Budget struct {
	Id string `json:id`
	Name string `json:name`
	Last_modified_on string `json:last_modified_on`
	First_month string `json:first_month`
	Last_month string `json:first_month`
	Date_format struct{
		Format string `json:format`
	} `json:date_format`
	Currency_format struct{
		IsoCode string `json:iso_code`
		ExampleFormat string `json:example_format`
		DecimalDigits int `json:decimal_digits`
		DecimalSeperator string `json:decimal_separator`
		GroupSeperator string `json:group_separator`
		CurrencySymbol string `json:currency_symbol`
		SymbolFirst bool `json:symbol_first`
		DisplaySymbol bool `json:display_symbol`

	} `json:currency_format`
}

type Budgets struct {
	Data struct {
		Budgets       []Budget `json:budgets`
		DefaultBudget Budget   `json:default_budget`
	}`json:"data"`
}


