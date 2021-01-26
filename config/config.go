package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Config struct {
	PersonalAccessToken string `json:"personal_access_token"`
	SyncDate time.Time `json:"sync_date"`
	HttpTimeout time.Duration `json:"http_timeout"`
}

func ReadConfig() Config {
	conf := Config{
		SyncDate:            time.Time{},
		HttpTimeout: 	  10*time.Second, // default values
	}
	home := os.Getenv("HOME")
	data, err := ioutil.ReadFile(home + "/.ynab-sync/config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &conf)
	return conf
}

func WriteConfig(conf Config)  {
	home := os.Getenv("HOME")

	data, err := json.Marshal( &conf)
	ioutil.WriteFile(home+"/.ynab-sync/config.json", data, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
}
