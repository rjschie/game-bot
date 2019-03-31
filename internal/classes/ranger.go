package classes

import gamebot "github.com/rjschie/gamebot/internal"

// NewRangerLibrary .
func NewRangerLibrary() []gamebot.Ability {
	var abilties []gamebot.Ability

	AddAbility(abilties, gamebot.Ability{
		Name:  "Sniper",
		Level: 1,
		Cost:  7,
	})

	AddAbility(abilties, gamebot.Ability{
		Name:    "Pet",
		Level:   1,
		Persist: true,
		Cost:    8,
	})

	AddAbility(abilties, gamebot.Ability{
		Name:    "Wisdom of the Wild",
		Level:   1,
		Persist: true,
		Cost:    15,
	})

	AddAbility(abilties, gamebot.Ability{
		Name:    "Heightened Senses",
		Level:   1,
		Persist: true,
		Cost:    5,
	})

	AddAbility(abilties, gamebot.Ability{
		Name:  "Disengage",
		Level: 1,
		Cost:  5,
	})

	AddAbility(abilties, gamebot.Ability{
		Name:    "Ammunition",
		Level:   1,
		Persist: true,
		Cost:    3,
	})

	return abilties
}
