package main

import (
	"database/sql"
	"fmt"

	"github.com/AVieraUY/go-forum/pkg/logger"

	"github.com/AVieraUY/go-forum/app"
	"github.com/AVieraUY/go-forum/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {
	c, err := app.InitApp(config.DevFilename)
	if err != nil {
		errors.Wrap(err, "InitApp")
	}

	ds := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", c.DBConfig.DbUser, c.DBConfig.DbPassword, c.DBConfig.DbHost, c.DBConfig.DbName)
	db, err := sql.Open("mysql", ds)
	if err != nil {
		logger.Log.Errorf("%s", err.Error())
	}

	defer db.Close()

}
