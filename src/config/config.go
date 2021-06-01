package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()

	config.SetConfigName(env)

	config.SetConfigType("yaml")
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	config.AddConfigPath("src/config/")

	config.AutomaticEnv()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = config.ReadInConfig()
	if err != nil {
		log.Fatalf("error on parsing configuration file %s", err)
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}

func GetSecretKey() string {
	configs := GetConfig()
	key := configs.GetString("security.secret_key")

	return key
}
