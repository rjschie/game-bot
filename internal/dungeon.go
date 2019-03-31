package gamebot

import (
	"math/rand"
	"time"

	"github.com/rjschie/gamebot/tabler"

	"github.com/fatih/color"
)

// Dungeon .
type Dungeon struct {
	Name      string
	Level     int
	Bosses    []Mob
	Mobs      []Mob
	Modifiers []Modifier

	CurrentBoss Mob
	DrawnMobs   []Mob
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
	mob, arr := d.Mobs[len(d.Mobs)-1], d.Mobs[:len(d.Mobs)-1]
	d.Mobs = arr
	d.DrawnMobs = append(d.DrawnMobs, mob)

	return mob
}

// DrawModifier .
func (d *Dungeon) DrawModifier() Modifier {
	modifier, arr := d.Modifiers[len(d.Modifiers)-1], d.Modifiers[:len(d.Modifiers)-1]
	d.Modifiers = arr

	if modifier.Type == "Encounter" {
		color.HiCyan("Encounter:\t%v\n\n", modifier.Name)
		modifier.UseAbility()
	}

	return modifier
}

// DrawMobWithModifier .
func (d *Dungeon) DrawMobWithModifier() Mob {
	mob := d.DrawMob()
	modifier := d.DrawModifier()

	if modifier.Type != "Encounter" {
		mob.Modifiers = append(mob.Modifiers, modifier)
	}

	return mob
}

// ShowDrawnMobs .
func (d Dungeon) ShowDrawnMobs() {
	var rows [][]string
	for _, mob := range d.DrawnMobs {
		rows = append(rows, []string{mob.Name})
	}

	tabler.RenderRows(
		nil,
		rows...,
	)
}

// ShowDeck .
func (d Dungeon) ShowDeck() {
	tabler.RenderColumns(
		[]string{"Mobs", "Modifiers"},
		tabler.Columnize(d.Mobs),
		tabler.Columnize(d.Modifiers),
	)
}
