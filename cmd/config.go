package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func setupConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	viper.SetConfigName("hithlain")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)

	viper.SetDefault(
		"libPath",
		filepath.Join(home, "/hitlain.db"),
	)
	viper.AutomaticEnv()

	return nil

}
