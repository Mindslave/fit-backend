package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Mindslave/fit-backend/internal/postgresql"
)

type server struct {
	Store *postgresql.Store
	Router *gin.Engine
}

func NewServer() *server {
	s := &server{}
	return s
}

func (s *server) Start(address string) {
	s.Router.Run(address)
}