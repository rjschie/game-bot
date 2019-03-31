package main

import (
	"fmt"

	"github.com/rjschie/gamebot/internal/env"
	"github.com/rjschie/gamebot/internal/models"
)

func main() {
	heroes := []models.Hero{
		models.NewHero("Player 1"),
		models.NewHero("Player 2"),
	}

	ctx := env.NewContext(heroes)

	fmt.Println(ctx.Heroes)
}
