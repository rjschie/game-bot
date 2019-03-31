package classes

import gamebot "github.com/rjschie/gamebot/internal"

// NewWarriorDeck .
func NewWarriorDeck(hero *gamebot.Hero) {
	AddAbility(hero, gamebot.Ability{
		Name:  "Aggro Lvl 1",
		Level: 1,
		Cost:  10,
	})

	AddAbility(hero, gamebot.Ability{
		Name:  "Sharpening Stone",
		Level: 1,
		Cost:  5,
	})

	AddAbility(hero, gamebot.Ability{
		Name:  "Slice",
		Level: 1,
		Cost:  5,
	})

	AddAbility(hero, gamebot.Ability{
		Name:    "Breastplate",
		Level:   1,
		Persist: true,
		Cost:    5,
	})

	AddAbility(hero, gamebot.Ability{
		Name:    "Shield",
		Level:   1,
		Persist: true,
		Cost:    8,
	})

	AddAbility(hero, gamebot.Ability{
		Name:    "Seasonsed",
		Level:   1,
		Persist: true,
		Cost:    5,
	})
}
