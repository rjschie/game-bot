package models

// Mob .
type Mob struct {
	Name      string
	Health    int
	Attack    int
	XP        int
	Modifiers []Modifier
	Ability   func()
}

// UseAbility .
func (m Mob) UseAbility() {
	if m.Ability != nil {
		m.Ability()
	}
}

func (m Mob) String() string {
	return m.Name
}
