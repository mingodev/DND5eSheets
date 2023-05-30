package characters

type Race struct {
	Name                string
	AttributesModifiers Attributes
	Speed               int
	Size                string
	Languages           []string
	// Feats
}

var (
	HUMAN = Race{
		Name: "Human",
		AttributesModifiers: Attributes{
			Strength:     Attribute{Score: 1},
			Dexterity:    Attribute{Score: 1},
			Constitution: Attribute{Score: 1},
			Intelligence: Attribute{Score: 1},
			Wisdom:       Attribute{Score: 1},
			Charisma:     Attribute{Score: 1},
		},
	}
	HUMAN_VARIANT = Race{
		Name: "Human Variant",
		AttributesModifiers: Attributes{
			Any: Attribute{Score: 2},
		},
	}
)
