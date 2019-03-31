package models

// Hero .
type Hero struct {
	Name   string
	Health int
	XP     int
	Attack int
	Defend int
	// Abilities []Ability
}

// NewHero .
func NewHero(name string) Hero {
	return Hero{
		Name:   name,
		Health: 10,
		XP:     10,
	}
}
