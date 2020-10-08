package fee

//go:generate stringer -type=state
type state int

const (
	_ state = iota
	UnderInspection
	Approved
	InAction
	Suspended
	UnderReturn
	End
)

type stateTransitions struct {
	allowed map[state][]state
}

func NewStateTransitions() *stateTransitions {
	return &stateTransitions{
		allowed: map[state][]state{
			UnderInspection: {
				UnderReturn,
				Approved,
			},
			Approved: {
				InAction,
				End,
			},
			InAction: {
				Suspended,
				End,
			},
			Suspended: {
				InAction,
				End,
			},
			UnderReturn: {
				UnderInspection,
				End,
			},
			End: {},
		},
	}
}

func (s *stateTransitions) IsAllowed(from, to state) bool {
	for _, allow := range s.allowed[from] {
		if allow == to {
			return true
		}
	}
	return false
}
