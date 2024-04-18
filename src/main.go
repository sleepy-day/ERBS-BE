package main

import (
	"fmt"
	"github.com/sleepy-day/ERBS-BE/src/util"
)

func main() {
	conf := util.GetDatabaseConfig()
	fmt.Println(conf.Database.Name)
}
