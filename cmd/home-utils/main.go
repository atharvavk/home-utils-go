package main

import "home-utils/internal/app"

func main() {
	appCtx := app.IntializeApp()
	app.CreateAndStartServer(appCtx)
}
