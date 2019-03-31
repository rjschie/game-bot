package dungeons

import gamebot "github.com/rjschie/gamebot/internal"

// AppendMob .
func AppendMob(d *gamebot.Dungeon, num int, m gamebot.Mob) *gamebot.Dungeon {
	for i := 0; i < num; i++ {
		d.Mobs = append(d.Mobs, m)
	}
	return d
}

// AppendModifier .
func AppendModifier(d *gamebot.Dungeon, num int, m gamebot.Modifier) *gamebot.Dungeon {
	for i := 0; i < num; i++ {
		d.Modifiers = append(d.Modifiers, m)
	}
	return d
}
