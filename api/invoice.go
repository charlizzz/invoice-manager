package api

import (
	"database/sql"
	"net/http"

	db "github.com/charlizzz/invoice-manager/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createInvoiceRequest struct {
	UserID int32   `json:"user_id" binding:"required"`
	Label  string  `json:"label" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

func (server *Server) createInvoice(ctx *gin.Context) {
	var req createInvoiceRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateInvoiceParams{
		UserID: sql.NullInt32{Int32: req.UserID, Valid: true},
		Label:  req.Label,
		Amount: int64(req.Amount),
	}

	_, err2 := server.store.CreateInvoice(ctx, arg)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err2))
		return
	}

	ctx.Status(http.StatusNoContent)
}
