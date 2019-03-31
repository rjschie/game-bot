package classes

import gamebot "github.com/rjschie/gamebot/internal"

// AddAbility .
func AddAbility(abilities []gamebot.Ability, ab gamebot.Ability) {
	for i := 0; i < 3; i++ {
		abilities = append(abilities, ab)
	}
}
