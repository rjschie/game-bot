package models

import (
	"fmt"
	"math/rand"
	"time"
)

// Enemy is
type Enemy struct {
	Name    string
	Attack  int
	Defense int
}

// Cull will return an enemy
func (e *Enemy) Generate(i int) Enemy {
	rand.Seed(time.Now().UnixNano())

	e.Name = fmt.Sprintf("Monster %v", i+1)
	e.Attack = rand.Intn(2) + 1
	e.Defense = rand.Intn(4) + 1
	return *e
}
