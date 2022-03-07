package _7_mediator

import "testing"

func TestMediator(t *testing.T) {
	lb := &listBox{}
	tb := &textButton{}
	mediator := &FormMediator{}

	mediator.lb = lb
	mediator.tb = tb

	lb.director = mediator
	tb.director = mediator

	mediator.ShowDialog()

	lb.Append(&textItem{"Hello", mediator})
	lb.Append(&textItem{"ListBox", mediator})

	mediator.ShowDialog()

	tb.Click()
	mediator.ShowDialog()

}
