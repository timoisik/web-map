package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/timoisik/web-map/cmd"
	"github.com/timoisik/web-map/models"
)

func main() {
	models.InitDB()
	defer models.Db.Close()

	cmd.Execute()
}
