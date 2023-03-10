package api

import (
	"net/http"

	db "github.com/charlizzz/invoice-manager/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (server *Server) listUsers(ctx *gin.Context) {
	//TODO: add query parameters request for a pagination
	// to be more safe in case of a huge list

	arg := db.ListUsersParams{
		Limit:  100,
		Offset: 0,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.IndentedJSON(http.StatusOK, users)
}
