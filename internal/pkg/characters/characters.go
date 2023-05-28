package characters

import (
	attributes "github.com/mingodev/DND5eSheets/internal/pkg/attributes"
	dice "github.com/mingodev/DND5eSheets/internal/pkg/dice"
	races "github.com/mingodev/DND5eSheets/internal/pkg/races"
	skills "github.com/mingodev/DND5eSheets/internal/pkg/skills"
)

type Character struct {
	Id               int // TODO : Find a better data type for UserID while implementing OAuth
	Race             races.Race
	Levels           []Level
	HitPointsCurrent int
	HitPointsMax     int
	HitDices         []dice.Die
	Attributes       attributes.Attributes
	Skills           skills.Skills
	Speed            int
}

type Level struct {
	HitPointsGained int
	// Class CharacterClass
	NewHitDie     dice.Die
	NewAttributes []attributes.Attributes
	// NewSpellSlots []SpellSlots
	// NewFeats []Feat
}
