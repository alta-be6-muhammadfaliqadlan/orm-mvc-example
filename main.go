package main

import (
	"immersive6/config"
	route "immersive6/routes"
)

func main() {
	config.InitDB()
	config.MigrateDBAuto()
	e := route.New()
	e.Start(":8000")
}
