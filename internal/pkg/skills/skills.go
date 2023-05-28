package skills

import "github.com/mingodev/DND5eSheets/internal/pkg/attributes"

type Skill struct {
	Name      string
	IsTrained bool
	IsExpert  bool
	Attribute attributes.Attribute
}

type Skills struct {
	Acrobatics     Skill
	AnimalHandling Skill
	Arcana         Skill
	Athletics      Skill
	Deception      Skill
	History        Skill
	Insight        Skill
	Intimidation   Skill
	Investigation  Skill
	Medicine       Skill
	Nature         Skill
	Perception     Skill
	Performance    Skill
	Persuasion     Skill
	Religion       Skill
	SleightOfHand  Skill
	Stealth        Skill
	Survival       Skill
}

var (
	ACROBATICS      = Skill{Name: "Acrobatics", Attribute: attributes.DEXTERITY}
	ANIMAL_HANDLING = Skill{Name: "Animal Handling", Attribute: attributes.WISDOM}
	ARCANA          = Skill{Name: "Arcana", Attribute: attributes.INTELLIGENCE}
	ATHLETICS       = Skill{Name: "Athletics", Attribute: attributes.STRENGTH}
	DECEPTION       = Skill{Name: "Deception", Attribute: attributes.CHARSIMA}
	HISTORY         = Skill{Name: "History", Attribute: attributes.INTELLIGENCE}
	INSIGHT         = Skill{Name: "Insight", Attribute: attributes.WISDOM}
	INTIMIDATION    = Skill{Name: "Intimidation", Attribute: attributes.CHARSIMA}
	INVESTIGATION   = Skill{Name: "Investigation", Attribute: attributes.INTELLIGENCE}
	MEDICINE        = Skill{Name: "Medicine", Attribute: attributes.WISDOM}
	NATURE          = Skill{Name: "Nature", Attribute: attributes.WISDOM}
	PERCEPTION      = Skill{Name: "Perception", Attribute: attributes.WISDOM}
	PERFORMANCE     = Skill{Name: "Performance", Attribute: attributes.CHARSIMA}
	PERSUASION      = Skill{Name: "Persuasion", Attribute: attributes.CHARSIMA}
	RELIGION        = Skill{Name: "Religion", Attribute: attributes.INTELLIGENCE}
	SLEIGHT_OF_HAND = Skill{Name: "Sleight of hand", Attribute: attributes.DEXTERITY}
	STEALTH         = Skill{Name: "Stealth", Attribute: attributes.DEXTERITY}
	SURVIVAL        = Skill{Name: "Survival", Attribute: attributes.WISDOM}
)
