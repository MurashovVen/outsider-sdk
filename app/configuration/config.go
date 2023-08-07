package configuration

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Default struct {
		Env Environment `desc:"(development)" default:"development" split_words:"true"`
	}

	GRPCServer struct {
		GRPCServerAddress string `desc:"host+port" default:"0.0.0.0:5000" split_words:"true"`
	}

	TelegramClient struct {
		TelegramBotToken string `desc:"Auth token" split_words:"true"`
	}

	Mongo struct {
		MongoURI string `default:"mongodb://mongo:27017" split_words:"true"`
	}

	RabbitMQ struct {
		RabbitMQURL string `default:"amqp://user:password@rabbitmq:5672/" split_words:"true"`
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
