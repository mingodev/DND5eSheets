package races

import "github.com/mingodev/DND5eSheets/internal/pkg/attributes"

type Race struct {
	Name                string
	AttributesModifiers attributes.Attributes
	Speed               int
	Size                string
	Languages           []string
	// Feats
}

var (
	// TODO : Implement races
	TEST_HUMAN = Race{
		Name: "Test Human",
		AttributesModifiers: attributes.Attributes{
			Strength:  attributes.Attribute{Name: attributes.STRENGTH.Name, Score: 1},
			Dexterity: attributes.Attribute{Name: attributes.DEXTERITY.Name, Score: 1},
		},
	}
)
