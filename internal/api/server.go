package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Mindslave/fit-backend/internal/postgresql"
)

type server struct {
	store *postgresql.Store
	router *gin.Engine
}

func NewServer() *server {
	s := &server{}
	router := gin.Default()
	s.router = router
	s.routes()
	return s
}