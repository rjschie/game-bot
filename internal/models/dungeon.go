package models

type Dungeon struct {
	Name      string
	Level     int
	Enemies   []Enemy
	Modifiers []Modifier
}
