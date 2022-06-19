package gorm

import (
	"github.com/go-playground/validator/v10"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	ConfigName                               string `json:"config_name" yaml:"config_name" toml:"config_name" validate:"required"`
	DSN                                      string `json:"dsn" yaml:"dsn" toml:"dsn" validate:"required,url"`
	LogLevel                                 int    `json:"log_level" yaml:"log_level" toml:"log_level" validate:"required"`
	DisableForeignKeyConstraintWhenMigrating bool   `json:"disable_foreign_key_constraint_when_migrating" yaml:"disable_foreign_key_constraint_when_migrating" toml:"disable_foreign_key_constraint_when_migrating"`
}

func (c *Config) GetConfigName() string {
	return c.ConfigName
}

func (c *Config) Validate() (err error) {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *Config) InitClient() (interface{}, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	logLevel := logger.LogLevel(c.LogLevel)
	return gorm.Open(mysql.Open(c.DSN), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logLevel),
		DisableForeignKeyConstraintWhenMigrating: c.DisableForeignKeyConstraintWhenMigrating,
	})
}
