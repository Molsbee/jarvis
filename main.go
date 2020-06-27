package main

import (
	"github.com/Molsbee/jarvis/cmd"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	cmd.Execute()
}
