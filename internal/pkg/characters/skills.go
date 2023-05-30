package characters

type Skill struct {
	Name      string
	IsTrained bool
	IsExpert  bool
	Attribute Attribute
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
	ACROBATICS      = Skill{Name: "Acrobatics", Attribute: DEXTERITY}
	ANIMAL_HANDLING = Skill{Name: "Animal Handling", Attribute: WISDOM}
	ARCANA          = Skill{Name: "Arcana", Attribute: INTELLIGENCE}
	ATHLETICS       = Skill{Name: "Athletics", Attribute: STRENGTH}
	DECEPTION       = Skill{Name: "Deception", Attribute: CHARSIMA}
	HISTORY         = Skill{Name: "History", Attribute: INTELLIGENCE}
	INSIGHT         = Skill{Name: "Insight", Attribute: WISDOM}
	INTIMIDATION    = Skill{Name: "Intimidation", Attribute: CHARSIMA}
	INVESTIGATION   = Skill{Name: "Investigation", Attribute: INTELLIGENCE}
	MEDICINE        = Skill{Name: "Medicine", Attribute: WISDOM}
	NATURE          = Skill{Name: "Nature", Attribute: WISDOM}
	PERCEPTION      = Skill{Name: "Perception", Attribute: WISDOM}
	PERFORMANCE     = Skill{Name: "Performance", Attribute: CHARSIMA}
	PERSUASION      = Skill{Name: "Persuasion", Attribute: CHARSIMA}
	RELIGION        = Skill{Name: "Religion", Attribute: INTELLIGENCE}
	SLEIGHT_OF_HAND = Skill{Name: "Sleight of hand", Attribute: DEXTERITY}
	STEALTH         = Skill{Name: "Stealth", Attribute: DEXTERITY}
	SURVIVAL        = Skill{Name: "Survival", Attribute: WISDOM}
)
