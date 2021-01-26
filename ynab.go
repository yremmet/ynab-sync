package main

import (
	"fmt"
	"github.com/yremmet/ynab/client"
	"github.com/yremmet/ynab/config"
	"log"
	"time"
)
func main() {
	conf := config.ReadConfig()
	fmt.Println(conf)
	c := client.NewYNABClient(conf)
	budgets, err := c.BudgetsService.GetBudgets()
	if (err != nil){
		handleError(err)
	}
	b := budgets.Data.Budgets[0]
	fmt.Println(b)
	accounts,err := c.AccountsService.GetAccounts(b)
	if (err != nil){
		handleError(err)
	}
	for _, a := range accounts.Data.Accounts {
		if (a.Note != ""){
			fmt.Printf("Exporting %s %s\n", a.Name, a.Note)
			ts := Export(a.Note, a.Id)
			err = c.TransactionsService.CreateTransactions(b, &ts)
			if (err != nil){
				handleError(err)
			}
		}
	}
	conf.SyncDate = time.Now()
	config.WriteConfig(conf)
}

func handleError(err *client.ClientError) {
	log.Fatal(err.Error())
}
