package cmd

import (
	"fmt"

	"github.com/rjschie/gamebot/internal/dungeons"
	"github.com/rjschie/gamebot/internal/models"

	tm "github.com/buger/goterm"
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

// App .
var App = grumble.New(&grumble.Config{
	Name:        "app",
	Description: "my app",
	Prompt:      "foo >> ",
	PromptColor: color.New(color.FgCyan, color.Bold),
})

func init() {
	swamp := dungeons.NewSwamp()

	App.AddCommand(&grumble.Command{
		Name: "deck",
		Help: "Shows dungeon deck",
		Run: func(c *grumble.Context) error {
			renderDeck(swamp)
			return nil
		},
	})

	drawCommand := &grumble.Command{
		Name: "draw",
		Help: "Draw card",
	}
	App.AddCommand(drawCommand)

	drawCommand.AddCommand(&grumble.Command{
		Name: "mob",
		Help: "Draw mob",
		Run: func(c *grumble.Context) error {
			fmt.Println(swamp.DrawMob().Name)
			return nil
		},
	})

	drawCommand.AddCommand(&grumble.Command{
		Name: "modifier",
		Help: "Draw modifier",
		Run: func(c *grumble.Context) error {
			fmt.Println(swamp.DrawModifier().Name)
			return nil
		},
	})

	drawCommand.AddCommand(&grumble.Command{
		Name: "mobmod",
		Help: "Draw mob with modifier",
		Run: func(c *grumble.Context) error {
			fmt.Println(swamp.DrawMobWithModifier().Name)
			return nil
		},
	})
}

func renderDeck(d *models.Dungeon) {
	table := tm.NewTable(0, 10, 5, ' ', 0)

	fmt.Fprintf(table, "Mobs\tModifiers\n")
	fmt.Fprintf(table, "--------\t--------\n")

	var ln int
	mobsLen := len(d.Mobs)
	modifiersLen := len(d.Modifiers)
	if ln = mobsLen; modifiersLen > ln {
		ln = modifiersLen
	}

	for i := 0; i < ln; i++ {
		var mobName string
		var modifierName string

		if i < mobsLen {
			mobName = d.Mobs[i].Name
		}

		if i < modifiersLen {
			modifierName = d.Modifiers[i].Name
		}

		fmt.Fprintf(table, "%s\t%s\n", mobName, modifierName)
	}

	tm.Println(table)

	tm.Flush()
}
