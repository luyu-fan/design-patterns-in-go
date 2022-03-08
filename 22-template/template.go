package _2_template

import "fmt"

type abstractClass struct {
}

func (ac *abstractClass) TemplateMethod() {
	// 算法的主调用流程
	fmt.Println("Stage 1")
	ac.PrimitiveMethod()
	// ...
	fmt.Println("Stage 2")
	ac.HookMethod()
}

// PrimitiveMethod 原语方法 在继承机制中应该实现为纯虚函数
func (ac *abstractClass) PrimitiveMethod() {}

// HookMethod 钩子方法 在继承机制中提供缺省实现
func (ac *abstractClass) HookMethod() {}

type concreteClass struct {
	*abstractClass
}

func (cc *concreteClass) PrimitiveMethod() {
	fmt.Println("Concrete Method")
}

func (cc *concreteClass) HookMethod() {
	fmt.Println("Concrete Hook Method")
}

func NewClass() *concreteClass {
	return &concreteClass{}
}
