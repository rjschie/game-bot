package dungeons

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
	gamebot "github.com/rjschie/gamebot/internal"
)

// NewSwamp .
func NewSwamp() gamebot.Dungeon {
	s := gamebot.Dungeon{}
	s.Name = "Swamp"
	s.Level = 1

	addBosses(&s)
	addMobs(&s)
	addModifiers(&s)

	s.ShuffleMobs()
	s.ShuffleModifiers()

	rand.Seed(time.Now().UTC().UnixNano())
	s.CurrentBoss = s.Bosses[rand.Intn(len(s.Bosses))]

	return s
}

func addBosses(s *gamebot.Dungeon) {
	s.Bosses = []gamebot.Mob{
		gamebot.Mob{
			Name:   "Alligator",
			Health: 20,
			Attack: 6,
			XP:     10,
		},
		gamebot.Mob{
			Name:   "Animated Slime",
			Health: 20,
			Attack: 4,
			XP:     10,
		},
		gamebot.Mob{
			Name:   "Boot Legger",
			Health: 20,
			Attack: 4,
			XP:     10,
		},
	}
}

func addMobs(s *gamebot.Dungeon) {
	AppendMob(s, 5, gamebot.Mob{
		Name:   "Cocka Roach",
		Health: 2,
		Attack: 1,
		XP:     1,
	})

	AppendMob(s, 2, gamebot.Mob{
		Name:   "Thieves",
		Health: 4,
		Attack: 3,
		XP:     3,
	})

	AppendMob(s, 2, gamebot.Mob{
		Name:   "Goblin",
		Health: 3,
		Attack: 3,
		XP:     2,
	})

	AppendMob(s, 3, gamebot.Mob{
		Name:   "Rats",
		Health: 2,
		Attack: 2,
		XP:     1,
	})

	AppendMob(s, 3, gamebot.Mob{
		Name:   "Spider",
		Health: 1,
		Attack: 3,
		XP:     1,
	})

	AppendMob(s, 2, gamebot.Mob{
		Name:   "Vagrant",
		Health: 4,
		Attack: 2,
		XP:     2,
	})
}

func addModifiers(s *gamebot.Dungeon) {
	AppendModifier(s, 5, gamebot.Modifier{
		Name: "Reinforcements",
		Type: "Encounter",
		Ability: func() {
			s.DrawMob()
		},
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Trip Wire",
		Type: "Encounter",
		Ability: func() {
			color.HiRed("You got tripped\n\n")
		},
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Spider Web",
		Type: "Encounter",
		Ability: func() {
			color.HiRed("You got webbed\n\n")
		},
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Spikes",
		Type: "Encounter",
		Ability: func() {
			color.HiRed("You got spiked\n\n")
		},
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Pile o' Poo",
		Type: "Encounter",
		Ability: func() {
			color.HiRed("You got poo'd\n\n")
		},
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Cave-in",
		Type: "Encounter",
		Ability: func() {
			color.HiRed("A Cave-in occurrs\n\n")
		},
	})

	AppendModifier(s, 3, gamebot.Modifier{
		Name: "Lucky Charm",
		Type: "Modifier",
	})

	AppendModifier(s, 3, gamebot.Modifier{
		Name: "Bone Armor",
		Type: "Modifier",
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Club",
		Type: "Modifier",
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Poison",
		Type: "Modifier",
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Net",
		Type: "Modifier",
	})

	AppendModifier(s, 1, gamebot.Modifier{
		Name: "Shiv",
		Type: "Modifier",
	})
}
