package _8_memento

import "testing"

func TestMemento(t *testing.T) {
	car := Car{}
	moveCmd := MoveCommand{
		target: car,
		memos:  make([]PosMemento, 0),
	}

	moveCmd.Execute()
	moveCmd.Execute()
	moveCmd.UnExecute()
	moveCmd.Execute()
	moveCmd.Execute()
	moveCmd.Execute()
	moveCmd.Execute()
	moveCmd.UnExecute()
	moveCmd.UnExecute()
	moveCmd.UnExecute()
	moveCmd.UnExecute()
	moveCmd.UnExecute()
	moveCmd.UnExecute()
	moveCmd.UnExecute()
	moveCmd.UnExecute()
}
