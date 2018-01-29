package game

type Stack struct {
	stack []Resolvable
}

func (s *Stack) PutOnTop(r Resolvable) {
	s.stack = append(s.stack, r)
}

func (s *Stack) ResolveAll(gs *GameState) {
	var r Resolvable

	for len(s.stack) > 0 {
		r, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
		r.Resolve(gs)
	}
}

func idGenerator() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type GameState struct {
	permanents []Permanent
	stack      Stack
}

type Card struct{}

type Zone struct {
	name  string
	owner string
	cards []Card
}

func (z *Zone) Append(card Card) {
	z.cards = append(z.cards, card)
}

type Permanent interface {
	MoveToZone(*Zone)
}

type Creature struct {
	name      string
	power     int
	toughness int
	card      Card
}

func (c Creature) MoveToZone(z *Zone) {
	z.Append(c.card)
}

type Resolvable interface {
	Resolve(*GameState)
}
