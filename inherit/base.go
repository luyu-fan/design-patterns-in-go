package inherit

import "fmt"

// Graphic 接口作为抽象类使用
type Graphic interface {
	Show()
}

// =========================================================================

// Shape 基本基类
type Shape struct {
	serialNumber int
}

func (s *Shape) GetSerialNumber() {
	fmt.Println("SerialNum:", s.serialNumber)
}

// Show 指针接收者
func (s *Shape) Show() {
	fmt.Println("Shape Showing! Num:", s.serialNumber)
}

// StructShow 值接收者
func (s Shape) StructShow() {
	fmt.Println("Shape Struct Showing! Num:", s.serialNumber)
}

// ModifyA 指针接收者
func (s *Shape) ModifyA(num int) {
	s.serialNumber = num
}

// ModifyB 值接收者
func (s Shape) ModifyB(num int) {
	s.serialNumber = num
}

// =========================================================================

type RigidShape struct {
	// 以内嵌结构体的方式完成对Shape的伪继承
	// 使用结构体声明和指针声明有什么区别?
	// A: 结构体继承->方法的重写指向关系混乱, 内外数值修改互不影响
	// B: 指针继承->方法的重写指向关系正常, 内外数值修改互相影响
	*Shape
	color int
}

func (r *RigidShape) GetColor() {
	fmt.Println("RigidShape Color:", r.color)
}

func (r *RigidShape) Show() {
	// 无论A还是B这里都能调用
	fmt.Println("RigidShape Showing!")
	r.Shape.Show()
}

// =========================================================================

type Rectangle struct {
	*RigidShape
	width  int
	height int
}

func (r *Rectangle) GetArea() {
	fmt.Println("Rectangle Area:", r.width*r.height)
}

func (r *Rectangle) GetPerimeter() {
	fmt.Println("Rectangle Perimeter:", 2*(r.width+r.height))
}

func (r *Rectangle) GetColor() {
	r.RigidShape.GetColor()
	fmt.Println("Rectangle Color:", r.color)
}

func (r *Rectangle) Show() {
	// 可以调用各个内嵌结构体的方法
	fmt.Println("Rectangle Show!")
	r.RigidShape.Show()
	r.Shape.Show()
}

//// Go不支持方法重载
//func (r *Rectangle) GetColor(value int) {
//	fmt.Println("Rectangle Color:", r.color)
//}
