package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	_ "github.com/lib/pq"

	"github.com/Mindslave/fit-backend/internal/api"
	"github.com/Mindslave/fit-backend/internal/config"
	"github.com/Mindslave/fit-backend/internal/postgresql"
	"github.com/Mindslave/fit-backend/internal/zap"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}	
}

func run() error {
	logger, err := zap.NewLogger()
	if err != nil {
		return errors.Wrap(err, "Could not instaniate logger")
	}
	config, err := config.Load("./env")
	if err != nil {
		return errors.Wrap(err, "Could not load config")
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return errors.Wrap(err, "Can't connect to database")
	}
	server := api.NewServer()
	server.Router = gin.New()
	server.Store = postgresql.NewStore(db)
	server.Routes()
	server.Start("127.0.0.1:8080")
	logger.Info("so far so good")
	return nil
}
