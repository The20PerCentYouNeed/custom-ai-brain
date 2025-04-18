package main

import (
	"github.com/The20PerCentYouNeed/custom-ai-brain/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
