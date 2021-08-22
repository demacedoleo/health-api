package main

import (
	app "github.com/demacedoleo/health-api/cmd/health"
	"github.com/demacedoleo/health-api/internal/platform/environment"
	"log"
	"os"
)

const (
	port = ":8080"
)

func main() {
	env := environment.GetFromString(os.Getenv("GO_ENVIRONMENT"))

	dependencies, err := app.BuildDependencies(env)
	if err != nil {
		log.Fatal("error at dependencies building", err)
	}

	app := app.Build(dependencies)
	if err := app.Run(port); err != nil {
		log.Fatal("error at app startup", err)
	}
}
