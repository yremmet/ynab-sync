package main

import (
	"encoding/csv"
	"fmt"
	"github.com/andybrewer/mack"
	"github.com/yremmet/ynab/client"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Export(iban string, accountId string) []client.Transaction {
	now := time.Now();
	end_date := now.Format("2006-01-02");
	date := time.Date(now.Year(),now.Month(),1,0,0,0,0, now.Location());
	start_date := date.Format("2006-01-02");


	exportFile, err := export(iban, start_date, end_date)
	if (err != nil){
		if (strings.Contains(err.Error(), "Locked database")){
			fmt.Println("Your MoneyMoney database seems locked. Unlock MoneyMoney and try again.")
			os.Exit(3)
		} else {
			fmt.Println(err.Error())
		}

	}
	return readCSV(exportFile, accountId)
}


func export(iban string, start_date string, end_date string) (string, error) {
	cmd := fmt.Sprintf("set result to export transactions from account \"%s\" from date \"%s\" to date \"%s\" as \"csv\"",
		iban, start_date, end_date)

	return  mack.Tell("MoneyMoney", cmd)
}

func readCSV(file string, id string) []client.Transaction {
	var transactions []client.Transaction
	f, err := os.Open(file)

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)
	r.Comma = ';'
	record, err := r.Read() // skip header
	for {

		record, err = r.Read()
		if err == io.EOF {
			break
		}
		amount := parseFloatByLocale(record[7], "DE-DE")
		date := formatDate(record[0])
		t := client.Transaction{
			Amount:       amount,
			Date:         date,
			Cleared:      client.TRANSACTION_CLEARED,
			CategoryName: record[2],
			Approved:     false,
			FlagColor:    client.TRANSACTION_FLAG_COLOR_PURPLE,
			AccountId:    id,
			PayeeName:    record[3],
			ImportId:     fmt.Sprintf("%s-%d",date,amount),
		}
		transactions = append(transactions, t)
	}
	return transactions
}

func formatDate(date string) string {
	parts := strings.Split(date,".")
	day, _ := strconv.ParseInt(parts[0],10,64)
	month, _ := strconv.ParseInt(parts[1],10,64)
	year, _ := strconv.ParseInt(parts[2],10,64)
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)

}

func parseFloatByLocale(amount string, locale string) int {
	if (locale == "DE-DE") {
		amount = strings.Replace(amount, ",", ".", -1)
		amount = strings.Replace(amount, ".", "", 1)
	}
	float, _ := strconv.ParseFloat(amount, 64)
	float *= -10

	return int(float)
}