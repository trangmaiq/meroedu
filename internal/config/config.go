package config

import (
	"github.com/spf13/viper"
)

type config struct {
	Filesystem struct {
		RelativePath string
	}
	Database struct {
		User                 string
		Password             string
		Host                 string
		Port                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
		}
	}
	Server struct {
		Address string
	}
}

// C ...
var C config

// ReadConfig ...
func ReadConfig(filepath string) error {
	Config := &C
	viper.SetConfigFile(filepath)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}
	return nil
}
