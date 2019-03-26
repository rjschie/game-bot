package models

// Mob .
type Mob struct {
	Name      string
	Health    int
	Attack    int
	XP        int
	Modifiers []Modifier
	// Ability func(m *Mob) *Mob
}

// UseAbility .
// func (m *Mob) UseAbility() *Mob {
// 	return m.Ability(m)
// }
