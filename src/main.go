package main

import (
	"fmt"
	"github.com/sleepy-day/ERBS-BE/src/erbsdb"
	"github.com/sleepy-day/ERBS-BE/src/grpcserver"
	"github.com/sleepy-day/ERBS-BE/src/util"
	"os"
)

func main() {
	var update bool
	var run bool

	if len(os.Args) == 1 {
		fmt.Println("No launch option provided, defaulting to run...")
		run = true
	} else {
		update = os.Args[1] == "update"
		run = os.Args[1] == "run"
	}

	initializeDatabase()

	if update {
		updateDatabase()
	} else if run {
		grpcserver.StartServer()
	}
}

func updateDatabase() {
	err := erbsdb.PopulateGameDataTables()
	if err != nil {
		panic(err)
	}

	erbsdb.UpdateLanguageFiles()
	file, err := os.Open("./lang/English.txt")
	if err != nil {
		file, err = os.Open("../lang/English.txt")
		if err != nil {
			panic(err)
		}
	}

	eng := util.ParseEnglishFile(file)
	err = erbsdb.PopulateTablesFromLangFile(eng)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully updated!")
}

func initializeDatabase() {
	conf := util.GetDatabaseConfig()
	connStr := util.GetConnectionString(conf)
	erbsdb.InitDb(connStr, conf.Api.Key)
}
