package characters

import (
	dice "github.com/mingodev/DND5eSheets/internal/pkg/dice"
)

type Class struct {
	Name         string
	HitDie       dice.Die
	Requirements Attributes
	// SpellSlots []characters.SpellSlots
}
