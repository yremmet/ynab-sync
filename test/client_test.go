package main

import (
	"github.com/stretchr/testify/assert"
	client "github.com/yremmet/ynab/client"
	"github.com/yremmet/ynab/config"
	"net/http"
	"testing"
)



func TestClient401(t *testing.T) {
	mc := &MockClient{}
	mc.DoFunc =  func(*http.Request) (*http.Response, error) {
		return &http.Response{
			Status:           "Unauthorized",
			StatusCode:       401,
			Body:             nil,
		}, nil
	}
	client := client.NewYNABClient(config.Config{})
	client.SetClient(mc)
	_, err := client.BudgetsService.GetBudgets()
	assert.Error(t,err,)
	assert.Equal(t, "401 Unauthorized", err.Error())
}