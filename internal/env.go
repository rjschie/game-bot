package gamebot

// Context .
type Context struct {
	CurrentDungeon *Dungeon
	Heroes         []Hero
}

// NewContext .
func NewContext(heroes []Hero, dungeon *Dungeon) *Context {
	ctx := new(Context)
	ctx.CurrentDungeon = dungeon
	ctx.Heroes = heroes

	return ctx
}
