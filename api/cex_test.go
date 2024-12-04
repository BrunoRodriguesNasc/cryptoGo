package api_test

import (
	"cripto/api"
	"testing"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	_, err = api.GetRate("BTC")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
