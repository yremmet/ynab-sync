package main

import (
	"encoding/json"
	"fmt"
	"github.com/yremmet/ynab/client"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Config struct {
	PersonalAccessToken string `toml:"personal_access_token"`
	SyncDate time.Time `toml:"sync_date"`
}

func readConfig() Config {
	var conf Config
	home := os.Getenv("HOME")
	data, err := ioutil.ReadFile(home + "/.ynab-sync/config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &conf)
	return conf
}

func writeConfig(conf Config)  {
	home := os.Getenv("HOME")

	data, err := json.Marshal( &conf)
	ioutil.WriteFile(home+"/.ynab-sync/config.json", data, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	conf := readConfig()
	fmt.Println(conf)
	c := client.NewYNABClient(conf.PersonalAccessToken)
	fmt.Println(c.BudgetsService.GetBudgets().Data.Budgets[0].Name)
	b := c.BudgetsService.GetBudgets().Data.Budgets[0]
	accounts := c.AccountsService.GetAccounts(b)

	for _, a := range accounts.Data.Accounts {
		if (a.Note != ""){
			fmt.Printf("Exporting %s %s\n", a.Name, a.Note)
			ts := Export(a.Note, a.Id)
			c.TransactionsService.CreateTransactions(b, &ts)
		}
	}
	conf.SyncDate = time.Now()
	writeConfig(conf)
}
