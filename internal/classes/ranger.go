package classes

import gamebot "github.com/rjschie/gamebot/internal"

// NewRangerDeck .
func NewRangerDeck(hero *gamebot.Hero) {
	AddAbility(hero, gamebot.Ability{
		Name:  "Sniper",
		Level: 1,
		Cost:  7,
	})

	AddAbility(hero, gamebot.Ability{
		Name:    "Pet",
		Level:   1,
		Persist: true,
		Cost:    8,
	})

	AddAbility(hero, gamebot.Ability{
		Name:    "Wisdom of the Wild",
		Level:   1,
		Persist: true,
		Cost:    15,
	})

	AddAbility(hero, gamebot.Ability{
		Name:    "Heightened Senses",
		Level:   1,
		Persist: true,
		Cost:    5,
	})

	AddAbility(hero, gamebot.Ability{
		Name:  "Disengage",
		Level: 1,
		Cost:  5,
	})

	AddAbility(hero, gamebot.Ability{
		Name:    "Ammunition",
		Level:   1,
		Persist: true,
		Cost:    3,
	})
}
