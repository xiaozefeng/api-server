package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"time"
)

var self *sql.DB

func Init() error {
	var username = viper.GetString("db.username")
	var passwd = viper.GetString("db.passwd")
	var host = viper.GetString("db.host")
	var port = viper.GetInt("db.port")
	var database = viper.GetString("db.database")
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, passwd, host, port, database)

	db, err := sql.Open("mysql", url)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	self = db
	return err
}

func GetDB() *sql.DB {
	return self
}
