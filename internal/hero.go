package gamebot

// Hero .
type Hero struct {
	Name    string
	Health  int
	XP      int
	Attack  int
	Defend  int
	Class   string
	Library []Ability
}

// NewHero .
func NewHero(name string) Hero {
	return Hero{
		Name: name,
	}
}
