package gamebot

// Hero .
type Hero struct {
	Name   string
	Health int
	XP     int
	Attack int
	Defend int
	Class  string

	Library []Ability
}
