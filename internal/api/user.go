package api

import (
	"net/http"

	"github.com/Mindslave/fit-backend/internal/postgresql"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	FirstName	string 	`json:"first_name" binding:"required"`
	LastName	string 	`json:"last_name" binding:"required"`
	Coach		bool	`json:"coach" binding:"required"`
	Goal		string	`json:"goal" binding:"required"`
}

func (s *server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	arg := postgresql.CreateUserParams{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Goal: req.Goal,
		Coach: req.Coach,
	}
	user, err := s.Store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)	
}