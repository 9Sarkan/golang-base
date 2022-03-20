package app

import (
	"github.com/9sarkan/golang-base/config"
)

type App struct {
}

// StartApplication function will start application and configure all dependencies
func (app App) StartApplication(dbg bool, cfgPath string) error {
	config.Map.SetDebugMode(dbg)
	if err := config.Map.LoadFromFile(cfgPath); err != nil {
		return err
	}
	return nil
}

// StartDatabases configure databases
func (app App) StartDatabases(cnf config.Database) error {
	return nil
}
