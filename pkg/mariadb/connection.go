package mariadb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	_ "oprec/pkg/config"
	"oprec/pkg/helper"
	"time"
)

func NewConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.Get("mariadb.username"),
		viper.Get("mariadb.password"),
		viper.Get("mariadb.host"),
		viper.Get("mariadb.port"),
		viper.Get("mariadb.database"),
	)

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(viper.GetInt("mariadb.pool_min"))
	db.SetMaxOpenConns(viper.GetInt("mariadb.pool_max"))
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
