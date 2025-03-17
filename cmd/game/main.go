package main

import (
	"number-guessing-game/internal/cli"
	"number-guessing-game/internal/repository"
	"number-guessing-game/internal/usecase"
)

func main() {
	// Initialize dependencies
	randomRepo := repository.NewDefaultRandomRepository()
	gameUsecase := usecase.NewGameUsecase(randomRepo)
	cli := cli.NewCLI(gameUsecase)

	// Run the CLI
	cli.Run()
}
