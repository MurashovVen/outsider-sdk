package rabbitmq

import (
	"github.com/streadway/amqp"

	"github.com/MurashovVen/outsider-sdk/app/logger"
)

type (
	Channel struct {
		*amqp.Channel

		logger *logger.Logger
	}

	ChannelOption func(*Channel)
)

func MustConnect(url string, opts ...ChannelOption) *Channel {
	ch, err := ChannelConnect(url, opts...)
	if err != nil {
		panic("can't connect rabbit: " + err.Error())
	}

	return ch
}

func ChannelConnect(url string, opts ...ChannelOption) (*Channel, error) {
	channel := &Channel{
		logger: logger.NewNop(),
	}

	for _, opt := range opts {
		opt(channel)
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel.Channel, err = conn.Channel()
	if err != nil {
		return nil, err
	}

	// todo
	//channel.Channel.NotifyReturn()

	return channel, nil
}

func ChannelWithLogger(log *logger.Logger) ChannelOption {
	return func(channel *Channel) {
		channel.logger = log.Named("RabbitChannel")
	}
}
