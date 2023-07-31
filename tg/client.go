package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	*tgbotapi.BotAPI
}

func MustCreateAndConnect(token string) *Client {
	client, err := CreateAndConnect(token)
	if err != nil {
		panic("connecting: " + err.Error())
	}

	return client
}

func CreateAndConnect(token string) (*Client, error) {
	tg, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Client{BotAPI: tg}, nil
}
