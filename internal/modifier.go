package gamebot

// Modifier .
type Modifier struct {
	Name    string
	Type    string
	Ability func()
}

// UseAbility .
func (m Modifier) UseAbility() {
	if m.Ability != nil {
		m.Ability()
	}
}

func (m Modifier) String() string {
	return m.Name
}
