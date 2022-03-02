package __bridge

import "fmt"

// Window 面向客户的核心接口
type Window interface {
	RenderWindow()
	RenderLabel()
	RenderAll()
	// 其它更多功能支撑性方法
	//Drag()
	//Move()
	//Resize()
}

// baseAbstractWindow 可以为空 也可以提供原始缺省功能实现 在Go里面使用结构体包裹
type baseAbstractWindow struct {
	impl WindowImpl
}

// basicWindow 任何平台下的一般性窗口 面向功能 可以有子类扩展
type basicWindow struct {
	*baseAbstractWindow
}

func (bw *basicWindow) RenderWindow() {
	fmt.Println("[BasicWindow RenderWindow]")
	bw.impl.DrawRect()
	bw.impl.DrawText()
}

func (bw *basicWindow) RenderLabel() {
	fmt.Println("[BasicWindow RenderLabel]")
	bw.impl.DrawLine()
	bw.impl.DrawLine()
	bw.impl.DrawLine()
	bw.impl.DrawLine()
	bw.impl.DrawText()
}

func (bw *basicWindow) RenderAll() {
	bw.RenderWindow()
	bw.RenderLabel()
}

// basicWindow 任何平台下的半透明窗口 面向功能 可以有子类扩展
type transparentWindow struct {
	*baseAbstractWindow
}

func (tw *transparentWindow) RenderWindow() {
	fmt.Println("[TransparentWindow RenderWindow]")
	fmt.Println("[Set Opacity]")
	tw.impl.DrawRect()
	tw.impl.DrawLine()
	tw.impl.DrawText()
}

func (tw *transparentWindow) RenderLabel() {
	fmt.Println("[TransparentWindow RenderLabel]")
	fmt.Println("[Set Opacity]")
	tw.impl.DrawLine()
	tw.impl.DrawLine()
	tw.impl.DrawText()
}

func (tw *transparentWindow) RenderAll() {
	tw.RenderWindow()
	tw.RenderLabel()
}

// WindowImpl 面向实现的接口封装 是一些较细粒度的基本实现
type WindowImpl interface {
	DrawRect()
	DrawLine()
	DrawText()
	// 其它更多基础方法实现
	//DrawCorner()
	//DrawColor()

}

// xWindowImpl XWindow平台下的实现
type xWindowImpl struct {
}

func (xwi *xWindowImpl) DrawRect() {
	fmt.Println("xWindowImpl is drawing a Rectangle!")
}

func (xwi *xWindowImpl) DrawLine() {
	fmt.Println("xWindowImpl is drawing a Line!")
}

func (xwi *xWindowImpl) DrawText() {
	fmt.Println("xWindowImpl is drawing a Text!")
}

// pMWindowImpl PMWindow平台下的实现
type pMWindowImpl struct {
}

func (pwi *pMWindowImpl) DrawRect() {
	fmt.Println("pMWindowImpl is drawing a Rectangle!")
}

func (pwi *pMWindowImpl) DrawLine() {
	fmt.Println("pMWindowImpl is drawing a Line!")
}

func (pwi *pMWindowImpl) DrawText() {
	fmt.Println("pMWindowImpl is drawing a Text!")
}

// NewWindow 窗口的构造函数 也是工厂方法
func NewWindow(wType string) Window {

	// 根据OS信息获取平台实现
	impl := &xWindowImpl{}

	if wType == "Basic" {
		return &basicWindow{&baseAbstractWindow{impl}}
	} else if wType == "Transparent" {
		return &transparentWindow{&baseAbstractWindow{impl}}
	} else {
		return &basicWindow{&baseAbstractWindow{impl}}
	}

}
