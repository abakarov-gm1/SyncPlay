package main

import (
	"goWin/services/syncPlay/internal/app"
	"goWin/services/syncPlay/internal/route"
)

func main() {
	services := app.NewApp()
	route.Server(services)
}
