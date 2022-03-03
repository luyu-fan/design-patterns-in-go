package __decorator

import "fmt"

// VisualComponent 视觉组件统一接口
type VisualComponent interface {
	Draw()
	Resize()
}

// textView 具体的被装饰类
type textView struct {
	text string
}

func (t *textView) Draw() {
	fmt.Println("TextView is drawing:", t.text)
}

func (t *textView) Resize() {
	fmt.Println("TextView is resizing:", t.text)
}

// visualDecorator 作为抽象装饰类
type visualDecorator struct {
	component VisualComponent
}

// Draw 缺省装饰方法 一般仅仅调用待装饰对象的同名方法
func (vd *visualDecorator) Draw() {
	fmt.Println("Default decorated drawing")
	vd.component.Draw()
}

// Resize 缺省装饰方法
func (vd *visualDecorator) Resize() {
	fmt.Println("Default decorated resizing")
	vd.component.Draw()
}

type borderDecorator struct {
	*visualDecorator
}

// Draw 装饰方法
func (bd *borderDecorator) Draw() {
	bd.setSize(12, 34)
	// 调用父类方法
	bd.visualDecorator.Draw()
}

// Resize 装饰方法
func (bd *borderDecorator) Resize() {
	bd.setSize(12, 34)
	// 调用父类方法
	bd.visualDecorator.Resize()
}

// setSize 自身私有方法 用于完成装饰工作
func (bd *borderDecorator) setSize(h, w int) {
	fmt.Println("Setting Size: ", h, w)
}
