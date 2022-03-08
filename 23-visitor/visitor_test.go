package _3_visitor

import "testing"

func TestVisitor(t *testing.T) {
	computer := NewComputer()
	visitor := NewVisitor()

	// 可以使用外部遍历等方式进行遍历
	computer.c.Accept(visitor)
	computer.d.Accept(visitor)
	computer.ch.Accept(visitor)
	computer.b.Accept(visitor)
}
