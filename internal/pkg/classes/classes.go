package classes

import (
	"github.com/mingodev/DND5eSheets/internal/pkg/attributes"
	dice "github.com/mingodev/DND5eSheets/internal/pkg/dice"
)

type Class struct {
	Name         string
	HitDie       dice.Die
	Requirements attributes.Attributes
	// SpellSlots []characters.SpellSlots
}
