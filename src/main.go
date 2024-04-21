package main

import (
	"github.com/sleepy-day/ERBS-BE/src/erbsdb"
	"github.com/sleepy-day/ERBS-BE/src/util"
)

func main() {
	conf := util.GetDatabaseConfig()
	connStr := util.GetConnectionString(conf)
	erbsdb.InitDb(connStr, conf.Api.Key)
	err := erbsdb.PopulateGameDataTables()
	if err != nil {
		panic(err)
	}
}
