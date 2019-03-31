package classes

import gamebot "github.com/rjschie/gamebot/internal"

// AddAbility .
func AddAbility(h *gamebot.Hero, ab gamebot.Ability) {
	for i := 0; i < 3; i++ {
		h.Library = append(h.Library, ab)
	}
}
