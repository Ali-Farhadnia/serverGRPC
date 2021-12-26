package cmd

import (
	"fmt"

	"github.com/Ali-Farhadnia/serverGRPC/config"
)

var AppConfig *config.AppConfig

func Config() error {
	AppConfig = config.NewAppConfig()

	/*
		parser := flags.NewParser(AppConfig, flags.Default)
		_, err := parser.Parse()
		if err != nil {
			return err
		}
	*/
	err := AppConfig.SetEnvVar()
	if err != nil {
		return err
	}
	fmt.Println(*AppConfig)
	return nil
}
