package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/atsanalyzer/config"
	"github.com/supermario64bit/atsanalyzer/routes"
)

func main() {
	config.LoadEnVFile()

	r := gin.Default()

	routes.MountHTTPRoutes(r)
	r.Run()
}
