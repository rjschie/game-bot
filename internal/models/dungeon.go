package models

import (
	"math/rand"
	"time"
)

// Dungeon .
type Dungeon struct {
	Name      string
	Level     int
	Bosses    []Mob
	Mobs      []Mob
	Modifiers []Modifier
}

// ShuffleMobs .
func (d *Dungeon) ShuffleMobs() {
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(d.Mobs), func(i, j int) {
		d.Mobs[i], d.Mobs[j] = d.Mobs[j], d.Mobs[i]
	})
}

// ShuffleModifiers .
func (d *Dungeon) ShuffleModifiers() {
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(d.Modifiers), func(i, j int) {
		d.Modifiers[i], d.Modifiers[j] = d.Modifiers[j], d.Modifiers[i]
	})
}

// DrawMob .
func (d *Dungeon) DrawMob() Mob {
	x, arr := d.Mobs[len(d.Mobs)-1], d.Mobs[:len(d.Mobs)-1]
	d.Mobs = arr

	return x
}

// DrawModifier .
func (d *Dungeon) DrawModifier() Modifier {
	x, arr := d.Modifiers[len(d.Modifiers)-1], d.Modifiers[:len(d.Modifiers)-1]
	d.Modifiers = arr

	return x
}

// DrawMobWithModifier .
func (d *Dungeon) DrawMobWithModifier() Mob {
	mob := d.DrawMob()
	mob.Modifiers = append(mob.Modifiers, d.DrawModifier())

	return mob
}
