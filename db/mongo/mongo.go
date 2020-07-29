package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Setting struct {
	Username string
	Password string
	Address  string
	Database string
	Options  map[string]string
}

func New(databaseSetting *Setting, debug bool) (*mongo.Client, error) {

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Address,
		databaseSetting.Database,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			return
		}
	}()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}
