package __composite

import "testing"

func TestComponent(t *testing.T) {

	p1 := &panel{}
	b1 := &button{}
	l1 := &label{}

	p2 := &panel{}
	b2 := &button{}
	l2 := &label{}

	p1.Add(b1)
	p1.Add(l1)
	p1.Add(p2)
	p2.Add(b2)
	p2.Add(l2)

	p1.Draw()

}
