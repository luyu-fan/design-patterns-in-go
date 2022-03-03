package _0_facade

import "fmt"

type Token interface{}

// scanner 内部文本扫描类
type scanner struct {
	stream interface{}
}

// Scan 扫描源码
func (s *scanner) Scan() Token {
	return nil
}

// parser 语法解析类
type parser struct {
}

// Parse 解析方法实现
func (p *parser) Parse(*scanner, *programBuilder) {
	fmt.Println("Parsing...")
}

// programBuilder 语法树构建类
type programBuilder struct {
	node *programNode
}

// NewVariable 返回变量结点
func (pb *programBuilder) NewVariable() *programNode {
	fmt.Println("NewVariable Node")
	return nil
}

// NewAssignment 返回赋值结点
func (pb *programBuilder) NewAssignment() *programNode {
	fmt.Println("NewAssignment Node")
	return nil
}

// NewStatement 返回语句结点
func (pb *programBuilder) NewStatement() *programNode {
	fmt.Println("Statement Node")
	return nil
}

// ... 其它实现

// programNode 结点类
type programNode struct {
}

// GetSourcePosition 获取源位置
func (pn *programNode) GetSourcePosition(line *int, index *int) {
	fmt.Println("GetSourcePosition:", line, index)
}

func (pn *programNode) Add(*programNode) {
	fmt.Println("Add Node")
	// ...
}

func (pn *programNode) Remove(*programNode) {
	fmt.Println("Remove Node")
	// ...
}

// codeGenerator 代码生成类
type codeGenerator struct {
	// ...
}

// Visit 遍历结点
func (cg *codeGenerator) Visit(node *programNode) {
	// ...
}

// Compiler 对外暴露的外观对象 也可以使用接口再次抽象封装
type Compiler struct {
}

// Compile 对外提供的接口 也是整个子系统的入口
func (c *Compiler) Compile(inputStream, outputStream interface{}) {
	// 将职责指派给编译子系统的内部的各个负责对象
	// ...
}

// NewDefaultCompiler 对外提供外观对象
func NewDefaultCompiler() *Compiler {
	return &Compiler{}
}
