package app

// This file was generated by the goro tool.
// Editing this file might prove futile when you re-run the goro commands

import (
	"log"

	"github.com/greyfox12/GoDilpom1/pkg/logger"
)

func (a *App) initLogger() {
	l, err := logger.NewLogger(a.cfg.App.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	a.logger = l

}
