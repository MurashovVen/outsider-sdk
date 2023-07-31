package configuration

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Default struct {
		Env Environment `desc:"(development)" default:"development" split_words:"true"`
	}

	GRPCServer struct {
		GRPCServerAddress string `desc:"host+port" default:"localhost:5000" split_words:"true"`
	}

	TelegramClient struct {
		TelegramBotToken string `desc:"Auth token" split_words:"true"`
	}
)

var configPrefix = "APP"

func MustProcessConfig(cfg interface{}) {
	if err := ProcessConfig(cfg); err != nil {
		panic("reading configuration: " + err.Error())
	}
}

func ProcessConfig(cfg interface{}) error {
	if err := envconfig.Usage(configPrefix, cfg); err != nil {
		return err
	}

	return envconfig.Process(configPrefix, cfg)
}
