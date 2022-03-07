package _9_observer

import "testing"

func TestObserver(t *testing.T) {

	sub := &ClockTimer{
		baseSubject{obs: map[signature]Observer{}},
	}

	obv := &DigitalClock{sig: 123, subject: sub}

	sub.Attach(obv)

	sub.Tick()
	sub.Tick()
	sub.Tick()
	sub.Tick()
}
