package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 4000
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "The amount cannot be greater than 1000", err.Error())

}

func TestTransactionWithAmountSmallerThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "The amount must be greater than 1", err.Error())

}
