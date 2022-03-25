package app

import (
	"context"
	"log"
	"time"

	"github.com/9sarkan/golang-base/config"
	"github.com/9sarkan/golang-base/pkg/middlewares"
	"github.com/9sarkan/golang-base/pkg/utils"
)

type App struct{}

// StartApplication function will start application and configure all dependencies
func (app App) StartApplication(dbg bool, cfgPath string) error {
	config.Map.SetDebugMode(dbg)
	if err := config.Map.LoadFromFile(cfgPath); err != nil {
		return err
	}
	log.Default().Println("config loaded.")
	if err := app.StartServers(); err != nil {
		return err
	}
	return nil
}

// StartDatabases configure databases
func (app App) StartDatabases(cnf config.Database) error {
	return nil
}

func (app App) StartServers() error {
	// start grpc server
	rc := RouterCenter{}
	if err := rc.InitialGrpcServer(context.Background(), config.Map.Service.GRPCServer, middlewares.GRPCAuth, utils.PanicRecover); err != nil {
		return err
	}
	time.Sleep(time.Hour)
	return nil
}
