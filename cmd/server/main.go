package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/Mindslave/fit-backend/internal/zap"
	"github.com/Mindslave/fit-backend/internal/config"
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
	logger.Info(config.DBDriver)
	return nil
}


func newDB() error {
	return nil
}