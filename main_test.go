package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLTPHandler(t *testing.T) {
	// Mock the latest prices
	prices["BTC/EUR"] = "45000.00"
	prices["BTC/USD"] = "50000.00"
	prices["BTC/CHF"] = "48000.00"

	expectedResult := `{"ltp":[{"pair":"BTC/EUR","amount":"45000.00"},{"pair":"BTC/USD","amount":"50000.00"},{"pair":"BTC/CHF","amount":"48000.00"}]}`

	req, err := http.NewRequest("GET", "/api/v1/ltp", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getLTPHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	assert.JSONEq(t, expectedResult, rr.Body.String(), "handler returned unexpected body")
}

func TestFetchTickers(t *testing.T) {
	// Set a timeout for the test
	timeout := time.After(10 * time.Second)
	done := make(chan bool)

	go func() {
		fetchTickers()
		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("Test timed out")
	case <-done:
		mu.RLock()
		defer mu.RUnlock()
		assert.NotEmpty(t, prices["BTC/EUR"], "BTC/EUR price should be set")
		assert.NotEmpty(t, prices["BTC/USD"], "BTC/USD price should be set")
		assert.NotEmpty(t, prices["BTC/CHF"], "BTC/CHF price should be set")
	}
}
