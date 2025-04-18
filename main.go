package main

import (
	"github.com/The20PerCentYouNeed/custom-ai-brain/db"
	"github.com/The20PerCentYouNeed/custom-ai-brain/routes"
)

func main() {
	db.InitDB()

	r := routes.SetupRouter()
	r.Run(":8080")
}
