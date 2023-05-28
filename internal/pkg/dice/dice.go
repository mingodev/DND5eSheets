package dice

type DieRollParameters struct {
	Advantage        bool `default:false`
	Disadvantage     bool `default:false`
	CanCritical      bool `default:false`
	CriticalTreshold int  `default:0`
}

type IRollable interface {
	Roll(params DieRollParameters) int
}

type Die struct {
	Faces int
}

func (die Die) Roll(params DieRollParameters) int {
	// TODO : Implement Roll Algorithm
	return die.Faces
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
