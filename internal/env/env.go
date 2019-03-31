package env

import (
	"github.com/rjschie/gamebot/internal/dungeons"
	"github.com/rjschie/gamebot/internal/models"
)

// Context .
type Context struct {
	CurrentDungeon *models.Dungeon
	Heroes         []models.Hero
}

// NewContext .
func NewContext(heroes []models.Hero) *Context {
	ctx := new(Context)
	ctx.CurrentDungeon = dungeons.NewSwamp()
	ctx.Heroes = heroes

	return ctx
}
