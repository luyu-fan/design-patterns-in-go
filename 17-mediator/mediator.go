package _7_mediator

import "fmt"

// DialogMediator 通用对话框显示中介者
type DialogMediator interface {
	// ShowDialog 显示
	ShowDialog()
	// WidgetChanged 组件发生改变时调用
	WidgetChanged(w Widget)
}

// Widget 通用组件接口
type Widget interface {
	// Changed 描述空间自身发生变化
	Changed()
	// HandleMouse 对鼠标交互作出响应
	HandleMouse(event interface{})
	// Display 打印展示
	Display()
	// 更多的用于获取内部信息的接口方法
	// ...
}

type listBox struct {
	director DialogMediator
	items    []Widget
}

func (lb *listBox) Changed() {
	// 自身变化时通知中介者进行进一步处理
	fmt.Println("list box is changed")
	lb.director.WidgetChanged(lb)
}

func (lb *listBox) HandleMouse(event interface{}) {
	// 响应鼠标事件
	fmt.Println("Fuck Keyboard!")
}

func (lb *listBox) Display() {
	fmt.Println("list box is displaying")
	for i := range lb.items {
		lb.items[i].Display()
	}
}

// Append 对象自身的处理方法
func (lb *listBox) Append(item Widget) {
	lb.items = append(lb.items, item)
	lb.Changed()
}

// TextItem 一般的文本控件 一般可经常和其它控件搭配使用 可以选择保持有中介则的引用
type textItem struct {
	text     string
	director DialogMediator
}

func (ti *textItem) Changed() {
	ti.director.WidgetChanged(ti)
}

func (ti *textItem) HandleMouse(event interface{}) {
	fmt.Println("Invalid Mouse Event")
}

func (ti *textItem) Display() {
	fmt.Println(ti.text)
}

type textButton struct {
	director DialogMediator
	text     string
}

func (tb *textButton) Changed() {
	tb.director.WidgetChanged(tb)
}

func (tb *textButton) HandleMouse(event interface{}) {
	fmt.Println("Handle Mouse Event", event)
	tb.text = event.(string)
	tb.Changed()
}

func (tb *textButton) Display() {
	fmt.Println("Text Button", tb.text)
}

func (tb *textButton) Click() {
	tb.HandleMouse("click")
}

type FormMediator struct {
	lb Widget
	tb Widget
}

func (fm *FormMediator) ShowDialog() {
	fm.lb.Display()
	fm.tb.Display()
}

func (fm *FormMediator) WidgetChanged(w Widget) {
	// 中介者内部定义好各个对象之间交互时的通信组织方式 当某一个对象发生改变时
	// 能够个根据这个对象的身份和种类获取对应的内部数据
	// 并用于其它对象的通信
	// 一般情况下往往某一个中介者只关心某一个功能的对象组织和通信
	switch w.(type) {
	case *listBox:
		{
			fmt.Println("ListBox is changed")
			fm.lb.Display()
		}
	case *textButton:
		{
			fmt.Println("TextButton is changed")
			fm.tb.Display()
		}
	default:
		fmt.Println("Ignore this Widget")
	}
}
