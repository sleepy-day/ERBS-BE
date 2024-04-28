package main

import (
	"fmt"
	"github.com/sleepy-day/ERBS-BE/src/erbsdb"
	"github.com/sleepy-day/ERBS-BE/src/util"
	"os"
)

func main() {
	update := os.Args[1] == "update"
	run := os.Args[1] == "run"

	if !update && !run {
		fmt.Println("No launch option provided, defaulting to run...")
		run = true
	}

	initializeDatabase()

	if update {
		updateDatabase()
	} else if run {
		fmt.Println("Not yet implemented.")
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
