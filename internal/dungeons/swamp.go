package dungeons

import (
	mdl "github.com/rjschie/gamebot/internal/models"
)

// NewSwamp .
func NewSwamp() *mdl.Dungeon {
	s := new(mdl.Dungeon)
	s.Name = "Swamp"
	s.Level = 1

	addBosses(s)
	addMobs(s)
	addModifiers(s)

	s.ShuffleMobs()
	s.ShuffleModifiers()

	return s
}

func addBosses(s *mdl.Dungeon) {
	s.Bosses = []mdl.Mob{
		mdl.Mob{
			Name:   "Alligator",
			Health: 20,
			Attack: 6,
			XP:     10,
		},
		mdl.Mob{
			Name:   "Animated Slime",
			Health: 20,
			Attack: 4,
			XP:     10,
		},
		mdl.Mob{
			Name:   "Boot Legger",
			Health: 20,
			Attack: 4,
			XP:     10,
		},
	}
}

func addMobs(s *mdl.Dungeon) {
	appendMob(s, 5, mdl.Mob{
		Name:   "Cocka Roach",
		Health: 2,
		Attack: 1,
		XP:     1,
	})

	appendMob(s, 2, mdl.Mob{
		Name:   "Thieves",
		Health: 4,
		Attack: 3,
		XP:     3,
	})

	appendMob(s, 2, mdl.Mob{
		Name:   "Goblin",
		Health: 3,
		Attack: 3,
		XP:     2,
	})

	appendMob(s, 3, mdl.Mob{
		Name:   "Rats",
		Health: 2,
		Attack: 2,
		XP:     1,
	})

	appendMob(s, 3, mdl.Mob{
		Name:   "Spider",
		Health: 1,
		Attack: 3,
		XP:     1,
	})

	appendMob(s, 2, mdl.Mob{
		Name:   "Vagrant",
		Health: 4,
		Attack: 2,
		XP:     2,
	})
}

func addModifiers(s *mdl.Dungeon) {
	appendModifier(s, 5, mdl.Modifier{
		Name: "Reinforcements",
		Type: "Encounter",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Trip Wire",
		Type: "Encounter",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Spider Web",
		Type: "Encounter",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Spikes",
		Type: "Encounter",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Pile o' Poo",
		Type: "Encounter",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Cave-in",
		Type: "Encounter",
	})

	appendModifier(s, 3, mdl.Modifier{
		Name: "Lucky Charm",
		Type: "Modifier",
	})

	appendModifier(s, 3, mdl.Modifier{
		Name: "Bone Armor",
		Type: "Modifier",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Club",
		Type: "Modifier",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Poison",
		Type: "Modifier",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Net",
		Type: "Modifier",
	})

	appendModifier(s, 1, mdl.Modifier{
		Name: "Shiv",
		Type: "Modifier",
	})
}

func appendMob(d *mdl.Dungeon, num int, m mdl.Mob) *mdl.Dungeon {
	for i := 0; i < num; i++ {
		d.Mobs = append(d.Mobs, m)
	}
	return d
}

func appendModifier(d *mdl.Dungeon, num int, m mdl.Modifier) *mdl.Dungeon {
	for i := 0; i < num; i++ {
		d.Modifiers = append(d.Modifiers, m)
	}
	return d
}
