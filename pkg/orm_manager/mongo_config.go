// Created at 11/29/2021 10:12 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package orm_manager

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Name   string
	DBName string
	DSN    string
}

func (m *MongoConfig) valid() (err error) {
	if m.Name == "" {
		err = fmt.Errorf("config name is empty")
		return
	}
	if m.DBName == "" {
		err = fmt.Errorf("config db name is empty")
		return
	}
	if m.DSN == "" {
		err = fmt.Errorf("config DNS is empty")
		return
	}
	return
}

func (m *MongoConfig) initClient() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.DSN))
	if err != nil {
		return nil, err
	}
	db := client.Database(m.DBName)
	return db, nil
}
