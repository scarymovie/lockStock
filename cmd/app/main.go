package main

import "lockStock/internal/app"

const configPath = "config/main"

func main() {
	app.Run(configPath)
}
