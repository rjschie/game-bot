package gamebot

// Mob .
type Mob struct {
	Name      string
	Health    int
	Attack    int
	XP        int
	Modifiers []Modifier
	Ability   func(*Context)
}

// UseAbility .
func (m Mob) UseAbility(ctx *Context) {
	if m.Ability != nil {
		m.Ability(ctx)
	}
}

func (m Mob) String() string {
	return m.Name
}
