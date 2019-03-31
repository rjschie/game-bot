package classes

import gamebot "github.com/rjschie/gamebot/internal"

// NewWarriorLibrary .
func NewWarriorLibrary() []gamebot.Ability {
	var abilities []gamebot.Ability

	AddAbility(abilities, gamebot.Ability{
		Name:  "Aggro Lvl 1",
		Level: 1,
		Cost:  10,
	})

	AddAbility(abilities, gamebot.Ability{
		Name:  "Sharpening Stone",
		Level: 1,
		Cost:  5,
	})

	AddAbility(abilities, gamebot.Ability{
		Name:  "Slice",
		Level: 1,
		Cost:  5,
	})

	AddAbility(abilities, gamebot.Ability{
		Name:    "Breastplate",
		Level:   1,
		Persist: true,
		Cost:    5,
	})

	AddAbility(abilities, gamebot.Ability{
		Name:    "Shield",
		Level:   1,
		Persist: true,
		Cost:    8,
	})

	AddAbility(abilities, gamebot.Ability{
		Name:    "Seasonsed",
		Level:   1,
		Persist: true,
		Cost:    5,
	})

	return abilities
}
