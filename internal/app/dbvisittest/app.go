package dbvisittest

import (
	"context"
	"db-visit-test/internal/pkg/dbhelper"
	"flag"
	"github.com/BurntSushi/toml"
	"path/filepath"
)

const CONFIG_PATH = "./config.toml"

func initConfig(user string, password string) *dbhelper.Config {
	cfg := &dbhelper.Config{}
	filePath, err := filepath.Abs(CONFIG_PATH)
	if err != nil {
		panic(err)
	}
	if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
		panic(err)
	}
	cfg.Database.User = user
	cfg.Database.Password = password
	return cfg
}

func Main() {
	var (
		user     string
		password string
	)

	flag.StringVar(&user, "u", "", "db user")
	flag.StringVar(&password, "p", "", "db password")
	flag.Parse()

	dbConfig := initConfig(user, password)

	dbHelper := dbhelper.New(context.Background(), &dbConfig.Database)
	dbHelper.PrintDBInfo()
}
