package main

import (
	"go-init-service/internal/commands"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	startLogo := figure.NewColorFigure("AVA-TOOL", "basic", "blue", false)
	startLogo.Print()
	copy := figure.NewColorFigure("AVA-TOOL 2024 - Alexy Van Den Abele", "term", "blue", true)
	copy.Print()

	// Appeler la fonction qui gère la création d'un nouveau service
	commands.CreateServiceCommand()
}
