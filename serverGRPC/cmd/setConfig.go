package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Ali-Farhadnia/serverGRPC/config"
)

var AppConfig *config.AppConfig

func Config() error {
	AppConfig = config.NewAppConfig()
	file, err := os.ReadFile("defult_config.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, AppConfig)
	if err != nil {
		return err
	}
	fmt.Println(*AppConfig)
	return nil
}
