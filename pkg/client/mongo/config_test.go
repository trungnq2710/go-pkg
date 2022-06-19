package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestConfig_GetConfigName(t *testing.T) {
	type fields struct {
		ConfigName string
		DBName     string
		DSN        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "success", fields: fields{ConfigName: "foo"}, want: "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				ConfigName: tt.fields.ConfigName,
				DBName:     tt.fields.DBName,
				DSN:        tt.fields.DSN,
			}
			if got := c.GetConfigName(); got != tt.want {
				t.Errorf("GetConfigName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_InitClient(t *testing.T) {
	type fields struct {
		ConfigName string
		DBName     string
		DSN        string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *mongo.Database
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				ConfigName: tt.fields.ConfigName,
				DBName:     tt.fields.DBName,
				DSN:        tt.fields.DSN,
			}
			got, err := c.InitClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		ConfigName string
		DBName     string
		DSN        string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "success", wantErr: false, fields: fields{ConfigName: "foo", DBName: "bar", DSN: "http://127.0.0.1:3000"}},
		{name: "ConfigName required", wantErr: true, fields: fields{ConfigName: "", DBName: "bar", DSN: "http://127.0.0.1:3000"}},
		{name: "DBName required", wantErr: true, fields: fields{ConfigName: "foo", DBName: "", DSN: "http://127.0.0.1:3000"}},
		{name: "DSN required", wantErr: true, fields: fields{ConfigName: "foo", DBName: "bar", DSN: ""}},
		{name: "DSN invalid", wantErr: true, fields: fields{ConfigName: "foo", DBName: "bar", DSN: "127.0.0.1:3000"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				ConfigName: tt.fields.ConfigName,
				DBName:     tt.fields.DBName,
				DSN:        tt.fields.DSN,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
