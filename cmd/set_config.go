package cmd

import (
	"github.com/Ali-Farhadnia/serverGRPC/config"
)

var AppConfig *config.AppConfig

func Config() error {
	AppConfig = config.NewAppConfig()
	err := AppConfig.SetEnvVar()
	if err != nil {
		return err
	}
	return nil
}
