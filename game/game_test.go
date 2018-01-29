package game

import "testing"

func TestResolveEmptyStack(t *testing.T) {
	stack := Stack{}
	gs := &GameState{stack: stack}

	stack.ResolveAll(gs)
}

type CreatureSpell struct{}

func (s CreatureSpell) Resolve(gs *GameState) {
	perm := Creature{"Goblin", 1, 1, Card{}}
	gs.permanents = append(gs.permanents, perm)
}

func TestResolveCreatureSpell(t *testing.T) {

	stack := Stack{}
	gs := &GameState{stack: stack}

	stack.ResolveAll(gs)

	spell := CreatureSpell{}
	stack.PutOnTop(spell)
	stack.ResolveAll(gs)

	if len(gs.permanents) != 1 {
		t.Error("GameState should have a permanent")
	}

	expected := Creature{"Goblin", 1, 1, Card{}}
	if gs.permanents[0] != expected {
		t.Error("Unexpected permanent")
	}
}

type DestroyCreatureSpell struct {
	target *Creature
}

func (spell DestroyCreatureSpell) Resolve(gs *GameState) {

}

func TestDestroyCreature(t *testing.T) {
	creature := Creature{"Goblin", 1, 1, Card{}}

	if 1 != 2 {
		t.Error("Error")
	}
}
