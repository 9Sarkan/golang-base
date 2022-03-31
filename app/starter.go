package app

import (
	"context"
	"fmt"
	"sync"

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
	return app.StartServers()
}

// StartDatabases configure databases
func (app App) StartDatabases(cnf config.Database) error {
	return nil
}

func (app App) StartServers() error {
	// start grpc server
	var wg sync.WaitGroup
	rc := RouterCenter{}

	wg.Add(1)
	go func() {
		if err := rc.InitialGrpcServer(context.Background(), config.Map.Service.GRPCServer, middlewares.GRPCAuth, utils.PanicRecover); err != nil {
			// call service down notifiers
			fmt.Printf("hello world")
			utils.ServiceDownNotifier(err, "<internal grpc service>")
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
