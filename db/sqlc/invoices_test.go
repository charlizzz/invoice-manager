package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestInvoice(t *testing.T) Invoices {
	arg := CreateInvoiceParams{
		UserID: sql.NullInt32{Int32: 3, Valid: true},
		Label:  "Test project",
		Amount: 542,
	}

	invoice, err := testQueries.CreateInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)

	require.Equal(t, arg.UserID, invoice.UserID)
	require.Equal(t, arg.Label, invoice.Label)
	require.Equal(t, arg.Amount, invoice.Amount)
	require.EqualValues(t, sql.NullString{String: "pending", Valid: true}, invoice.Status)

	require.NotZero(t, invoice.ID)

	return invoice
}

func TestCreateInvoice(t *testing.T) {
	createTestInvoice(t)
}

func TestGetInvoice(t *testing.T) {
	invoice1 := createTestInvoice(t)

	invoice2, err := testQueries.GetInvoice(context.Background(), invoice1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, invoice2)

	require.Equal(t, invoice1.UserID, invoice2.UserID)
	require.Equal(t, invoice1.Label, invoice2.Label)
	require.Equal(t, invoice1.Amount, invoice2.Amount)
	require.Equal(t, sql.NullString{String: "pending", Valid: true}, invoice2.Status)
}

func TestListInvoices(t *testing.T) {
	var limit int = 2
	for i := 0; i < 10; i++ {
		createTestInvoice(t)
	}

	arg := ListInvoicesParams{
		Limit:  int32(limit),
		Offset: 2,
	}

	invoices, err := testQueries.ListInvoices(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, invoices, limit)

	for _, invoice := range invoices {
		require.NotEmpty(t, invoice)
	}
}

func TestDeleteInvoice(t *testing.T) {
	invoice1 := createTestInvoice(t)
	err := testQueries.DeleteInvoice(context.Background(), invoice1.ID)

	require.NoError(t, err)

	invoice2, err := testQueries.GetInvoice(context.Background(), invoice1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, invoice2)
}
