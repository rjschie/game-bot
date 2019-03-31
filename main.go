package main

import (
	gamebot "github.com/rjschie/gamebot/internal"
	"github.com/rjschie/gamebot/internal/classes"
	"github.com/rjschie/gamebot/internal/dungeons"
)

func main() {
	heroes := []gamebot.Hero{
		gamebot.Hero{
			Name:    "Player 1",
			Health:  10,
			XP:      10,
			Class:   "Warrior",
			Library: classes.NewWarriorLibrary(),
		},
	}

	dungeons := []gamebot.Dungeon{
		dungeons.NewSwamp(),
	}

	ctx := gamebot.NewContext(heroes, dungeons)

	ctx.CurrentDungeon.ShowDeck()
}
