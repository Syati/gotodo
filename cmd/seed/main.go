package main

import (
	"database/sql"
	"fmt"
	"github.com/Syati/gotodo/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var opts struct {
	Path string `long:"path" description:"A name" required:"true"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Fatal(err)
	}

	c := config.GetConfig()
	con, err := sql.Open("mysql", c.GetString("db.dsn"))
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	defer con.Close()

	for _, path := range dirwalk(opts.Path) {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		result, err := con.Exec(string(b))
		if err != nil {
			log.Fatal(err)
		}
		num, _ := result.RowsAffected()
		fmt.Printf("Affected row: %d\n", num)
	}

}

func dirwalk(seedDir string) []string {
	seedDir, _ = filepath.Abs(seedDir)
	files, err := ioutil.ReadDir(seedDir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() && file.Name() == config.GetConfig().GetString("APP_ENV") {
			paths = append(paths, dirwalk(filepath.Join(seedDir, file.Name()))...)
			continue
		}
		path := filepath.Join(seedDir, file.Name())
		fmt.Println(path)
		paths = append(paths, path)
	}

	return paths
}
