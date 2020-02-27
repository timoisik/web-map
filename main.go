package main

import (
	"github.com/timoisik/web-map/cmd"
	"github.com/timoisik/web-map/models"
)

func main() {
	models.InitDB()
	cmd.Execute()
}
