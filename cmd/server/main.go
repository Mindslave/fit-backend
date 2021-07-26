package main

import (
	"fmt"
	"os"

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
		return err
	}
	logger.Info("I am Working")
	config, err := config.Load("./env")
	if err != nil {
		return err
	}
	logger.Info("config loaded succesfully, driver: ", config.DBDriver)
	return nil
}


func newDB() error {
	return nil
}