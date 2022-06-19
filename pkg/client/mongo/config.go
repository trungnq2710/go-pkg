package mongo

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ConfigName string `json:"config_name" yaml:"config_name" toml:"config_name" validate:"required"`
	DBName     string `json:"db_name" yaml:"db_name" toml:"db_name" validate:"required"`
	DSN        string `json:"dsn" yaml:"dsn" toml:"dsn" validate:"required,url"`
}

func (c *Config) GetConfigName() string {
	return c.ConfigName
}

func (c *Config) Validate() (err error) {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *Config) InitClient() (interface{}, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.DSN))
	if err != nil {
		return nil, err
	}
	db := client.Database(c.DBName)
	return db, nil
}
