package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransactionTx(t *testing.T) {
	store := NewStore(testDB)

	user := createTestUser(t)
	invoice := createTestInvoice(t)

	//TODO : add a loop to test concurrency,and the good behavior of postgres transactions
	err := store.TransactionTx(context.Background(), TransactionTxParams{
		InvoiceID:  invoice.ID,
		NewBalance: 654321,
		UserID:     user.ID,
	})

	require.NoError(t, err)
}
