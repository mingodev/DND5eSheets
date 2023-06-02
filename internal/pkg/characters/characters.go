package characters

import (
	dice "github.com/mingodev/DND5eSheets/internal/pkg/dice"
)

type Character struct {
	Id               int // TODO : Find a better data type for UserID while implementing OAuth
	Name             string
	Race             Race
	Levels           []Level
	HitPointsCurrent int
	HitPointsMax     int
	HitDices         []dice.Die
	Attributes       Attributes
	Skills           Skills
	Speed            int
}

type Level struct {
	HitPointsGained int
	// Class CharacterClass
	NewHitDie     dice.Die
	NewAttributes []Attributes
	// NewSpellSlots []SpellSlots
	// NewFeats []Feat
}
