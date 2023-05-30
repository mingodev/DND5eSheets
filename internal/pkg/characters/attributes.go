package characters

type Attribute struct {
	Name  string
	Score int
}

type Attributes struct {
	Any          Attribute
	Strength     Attribute
	Dexterity    Attribute
	Constitution Attribute
	Intelligence Attribute
	Wisdom       Attribute
	Charisma     Attribute
}

var (
	ANY          = Attribute{Name: "Any"}
	STRENGTH     = Attribute{Name: "Strength"}
	DEXTERITY    = Attribute{Name: "Dexterity"}
	CONSTITUTION = Attribute{Name: "Constitution"}
	INTELLIGENCE = Attribute{Name: "Intelligence"}
	WISDOM       = Attribute{Name: "Wisdom"}
	CHARSIMA     = Attribute{Name: "Charisma"}
)
