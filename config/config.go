package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config global singleton for configs
var Config ServerConfig

// ServerConfig root config struct
type ServerConfig struct {
	Instance *viper.Viper `json:"-"`
	App      App          `json:"app"`
	Server   Server       `json:"server"`
}

// App general configs
type App struct {
	Name string `json:"name"`
}

// Server server config struct
type Server struct {
	Port int `json:"port"`
}

// InitConfigs reads the config file into a global singleton
func InitConfigs() {
	runtimeViper := viper.New()

	runtimeViper.SetConfigType("toml")
	runtimeViper.AddConfigPath("./config")

	// read in general configs shared between all environments
	runtimeViper.SetConfigName("shared")
	runtimeViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	runtimeViper.AutomaticEnv()
	err := runtimeViper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found at %s. Shutting down.\n", err)
		} else {
			log.Fatalf("Failed to read config file: %s. Shutting down.\n", err)
		}
		return
	}

	env := os.Getenv("ENV")

	if env == "" {
		env = "dev"
	}

	runtimeViper.SetConfigName(env)

	// read in environment specific configs
	err = runtimeViper.MergeInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found at %s. Shutting down.\n", err)
		} else {
			log.Fatalf("Failed to read config file: %s. Shutting down.\n", err)
		}
		return
	}

	err = runtimeViper.Unmarshal(&Config)
	if err != nil {
		log.Fatal(err)
	}

	return
}
