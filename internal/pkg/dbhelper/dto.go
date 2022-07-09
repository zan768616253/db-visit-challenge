package dbhelper

import "context"

type DbHelper struct {
	ctx     context.Context
	connUrl string
}

type Config struct {
	Database DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int32
	DBName   string
}

type DBInfo struct {
	Name string
	Size int64
}
