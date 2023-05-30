package dice

import "math/rand"

type DieRollParameters struct {
	advantage    bool `default:"false"`
	disadvantage bool `default:"false"`
}

type IRollable interface {
	Roll(params DieRollParameters) int
	GetFaces() int
}

type Die struct {
	faces int
}

func (die Die) Roll(params DieRollParameters) int {
	rollResult := rand.Intn(die.faces) + 1

	if params.advantage == params.disadvantage {
		return rollResult
	}

	secondRollResult := die.Roll(DieRollParameters{})

	if (params.advantage && secondRollResult > rollResult) || (params.disadvantage && secondRollResult < rollResult) {
		rollResult = secondRollResult
	}

	return rollResult
}

func (die Die) GetFaces() int {
	return die.faces
}

var (
	COIN = Die{2}
	D4   = Die{4}
	D6   = Die{6}
	D8   = Die{8}
	D10  = Die{10}
	D12  = Die{12}
	D20  = Die{20}
	D100 = Die{100}
)
