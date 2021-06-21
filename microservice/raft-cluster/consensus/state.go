package consensus

type state string

const (
	fir state = "first"
	sec       = "second"
	thr       = "third"
)

var allowedState map[state][]state

func init() {
	allowedState = make(map[state][]state)
	allowedState[fir] = []state{sec, thr}
	allowedState[sec] = []state{thr}
	allowedState[thr] = []state{fir}
}

func (s *state) CanTransition(next state) bool {
	for _, n := range allowedState[*s] {
		if n == next {
			return true
		}
	}
	return false
}

func (s *state) Transition(next state) {
	if s.CanTransition(next) {
		*s = next
	}
}
