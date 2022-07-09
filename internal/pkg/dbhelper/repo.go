package dbhelper

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Repository = *DbHelper

// New Create DbHelper instance
func New(ctx context.Context, dbConfig *DBConfig) *DbHelper {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	return &DbHelper{
		connUrl: connStr,
		ctx:     ctx,
	}
}

// PrintDBInfo Print DB name and size
func (r *DbHelper) PrintDBInfo() {
	dbpool, err := pgxpool.Connect(r.ctx, r.connUrl)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(r.ctx, SQL_GETDBINFO)
	if err != nil {
		log.Printf("Unable to query to database: %v\n", err)
		panic(err)
	}
	defer rows.Close()

	var dbInfos []DBInfo
	for rows.Next() {
		var r DBInfo
		err = rows.Scan(&r.Name, &r.Size)
		if err != nil {
			log.Printf("Unable to retrieve value from row: %v\n", err)
			panic(err)
		}
		log.Printf("database: %v, size: %dMB\n", r.Name, r.Size)
		dbInfos = append(dbInfos, r)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Get error when doing query to database: %v\n", err)
		panic(err)
	}
}
