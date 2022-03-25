package config

import (
	"os"

	"github.com/spf13/viper"
)

var Map ConfigMap = ConfigMap{}

type ConfigMap struct {
	DebugMode       bool            `yaml:"-"`
	Service         Service         `yaml:"service"`
	Databases       Database        `yaml:"databases"`
	AuthCreditional AuthCreditional `yaml:"authCreditional"`
}

func (cnf *ConfigMap) SetDebugMode(mode bool) {
	cnf.DebugMode = mode
}

func (cnf *ConfigMap) GetDebugMode() bool {
	return cnf.DebugMode
}

func (cnf *ConfigMap) LoadFromFile(cfgPath string) error {
	if _, err := os.Stat(cfgPath); err != nil {
		return err
	}
	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&Map)
}
