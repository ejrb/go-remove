package target

type Permanent struct {}

type Player struct {}

type Targeter interface {
  Target(p Permanent)
}

type Spell struct {
  name string
  text string
  owner Player
  effect() bool
}

type TargetingSpell struct {
  *Spell
  target
}

func (ts *TargetingSpell) Target(p Permanent) {
  ts.target = p
}
