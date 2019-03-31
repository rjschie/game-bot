package main

import (
	gamebot "github.com/rjschie/gamebot/internal"
	"github.com/rjschie/gamebot/internal/dungeons"
)

func main() {
	heroes := []gamebot.Hero{
		gamebot.NewHero("Player 1"),
	}
	dungeon := dungeons.NewSwamp()

	ctx := gamebot.NewContext(heroes, dungeon)

	ctx.CurrentDungeon.ShowDeck()
}
