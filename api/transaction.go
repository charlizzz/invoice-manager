package api

import (
	"database/sql"
	"net/http"

	db "github.com/charlizzz/invoice-manager/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transactionRequest struct {
	InvoiceID int32   `json:"invoice_id" binding:"required,min=1"`
	Amount    float64 `json:"amount" binding:"required,gt=0"`
	Reference string  `json:"reference"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req transactionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	invoice, err := server.store.GetInvoice(ctx, req.InvoiceID)
	if !validInvoice(ctx, req, invoice, err) {
		return
	}

	user, err := server.store.GetUser(ctx, invoice.UserID.Int32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	newBalance := calculateNewBalance(invoice.Amount, user.Balance)

	arg := db.TransactionTxParams{
		InvoiceID:  req.InvoiceID,
		NewBalance: newBalance,
		UserID:     user.ID,
	}

	err = server.store.TransactionTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func validInvoice(ctx *gin.Context, req transactionRequest, invoice db.Invoices, err error) bool {
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if invoice.Amount != int64(req.Amount) {
		ctx.JSON(http.StatusBadRequest, struct {
			Err string
		}{Err: "error: The amount does not match with this invoice"})
		return false
	}

	if invoice.Status.String == "paid" {
		ctx.JSON(http.StatusUnprocessableEntity, struct {
			Err string
		}{Err: "error: The bill is already paid"})
		return false
	}
	return true
}

func calculateNewBalance(invoiceAmount int64, userBalance int64) int64 {
	return invoiceAmount + userBalance
}
