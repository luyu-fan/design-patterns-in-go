package _4_command

import "fmt"

// Command 命令接口
type Command interface {
	Execute()
}

// CMDReceiver 更抽象的接收者接口 在C++中可以用模板类实现简单的接收者
type CMDReceiver interface {
	Operation()
}

// openCommand 简单的具体命令封装
type openCommand struct {
	// 内部维持对接受者实例的引用和状态
	// ... states
	receiver CMDReceiver
}

func (oc *openCommand) Execute() {
	// 调用命令自身的处理流程或接收者的方法
	fmt.Println("OpenCommand is executing")
	oc.Do()
	oc.receiver.Operation()
}

func (oc *openCommand) Do() {
	fmt.Println("OpenCommand is doing things")
}

// commonReceiver 通用接收者
type commonReceiver struct {
}

func (cr *commonReceiver) Operation() {
	fmt.Println("Common Receiver is doing something")
}

// Invoker 一般请求调用者的封装 例如框架中的回调注册者
type Invoker struct {
	cmd Command
}

func (iv *Invoker) Call() {
	iv.cmd.Execute()
}

func NewOpenCommand(rec CMDReceiver) Command {
	return &openCommand{rec}
}

func NewCommonReceiver() CMDReceiver {
	return &commonReceiver{}
}
