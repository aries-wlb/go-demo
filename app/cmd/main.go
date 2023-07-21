package main

import (
	_ "github.com/joho/godotenv/autoload"
	"dipont.com/demo/app"
)

func main() {
	webApp,_,_ :=app.BuildApp()
	webApp.Run()
}
