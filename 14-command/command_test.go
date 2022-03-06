package _4_command

import "testing"

func TestCommand(t *testing.T) {
	// 客户创建并在调用者哪里注册
	receiver := NewCommonReceiver()
	opCmd := NewOpenCommand(receiver)
	invoker := Invoker{opCmd}

	//面向过程式的语言则可以将命令作为Call的参数传递
	invoker.Call()
}
