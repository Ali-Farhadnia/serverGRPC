package cmd

import (
	"fmt"

	"github.com/Ali-Farhadnia/serverGRPC/config"
	flags "github.com/jessevdk/go-flags"
)

var AppConfig *config.AppConfig

func Config() error {
	AppConfig = config.NewAppConfig()
	parser := flags.NewParser(AppConfig, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		return err
	}
	fmt.Println(*AppConfig)
	return nil
}
