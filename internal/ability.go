package gamebot

// Ability .
type Ability struct {
	Name    string
	Cost    int
	Level   int
	Persist bool
	Ability func(*Hero, *Context)
}
