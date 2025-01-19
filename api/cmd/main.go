package main

import "ctf-challenge/internal/app"

func main() {
	server := app.SetupRouter()
	server.Run(":5001")
}
