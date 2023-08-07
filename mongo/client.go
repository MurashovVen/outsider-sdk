package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/MurashovVen/outsider-sdk/app/logger"
)

type Client struct {
	*mongo.Client

	log *logger.Logger
}

func MustConnect(ctx context.Context, address string, logger *logger.Logger) *Client {
	cl, err := Connect(ctx, address, logger)
	if err != nil {
		panic("connecting mongo: " + err.Error())
	}

	return cl
}

func Connect(ctx context.Context, address string, logger *logger.Logger) (*Client, error) {
	cl, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI(address).
			SetConnectTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}

	logger = logger.Named("MongoClient")

	if err := cl.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping mongo: %w", err)
	}

	logger.Info("connected")

	return &Client{
		Client: cl,
		log:    logger,
	}, nil
}
