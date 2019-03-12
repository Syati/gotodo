package main

import (
	"fmt"
	"github.com/Syati/gotodo/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	"path"
	"runtime"
)

func main() {
	c := config.GetConfig()

	p := &mysql.Mysql{}
	d, err := p.Open("mysql://" + c.Get("db.dsn").(string))
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer d.Close()

	m, err := migrate.NewWithDatabaseInstance(
		migrateFilePath(),
		"gotodo",
		d,
	)
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	m.Up()
}

func migrateFilePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return "file://" + path.Join(path.Dir(filename), "../db/migrate")
}
