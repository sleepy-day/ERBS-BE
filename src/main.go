package main

import (
	"encoding/json"
	"fmt"
	"github.com/sleepy-day/ERBS-BE/src/erbsdb"
	"github.com/sleepy-day/ERBS-BE/src/util"
	"os"
)

func main() {
	conf := util.GetDatabaseConfig()
	connStr := util.GetConnectionString(conf)
	erbsdb.InitDb(connStr, conf.Api.Key)
	if false {
		err := erbsdb.PopulateGameDataTables()
		if err != nil {
			panic(err)
		}
	}
	file, err := os.Open("./lang/English.txt")
	if err != nil {
		panic(err)
	}
	eng := util.ParseEnglishFile(file)
	for k, v := range eng.CharacterInfo {
		fmt.Println(k)
		jsonStr, _ := json.Marshal(v)
		fmt.Println(string(jsonStr))
	}
}
