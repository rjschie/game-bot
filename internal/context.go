package gamebot

// Context .
type Context struct {
	CurrentDungeon *Dungeon
	Dungeons       []Dungeon
	Heroes         []Hero
}

// NewContext .
func NewContext(heroes []Hero, dungeons []Dungeon) *Context {
	ctx := new(Context)

	ctx.Dungeons = dungeons
	ctx.Heroes = heroes

	ctx.CurrentDungeon = &dungeons[0]

	return ctx
}
