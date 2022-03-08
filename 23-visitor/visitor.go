package _3_visitor

import "fmt"

// EquipmentVisitor 在定义之初需要确定访问的各个类 很少有新的类出现
type EquipmentVisitor interface {
	// VisitCPU 访问CPU具体类 因为已经知道了具体类则可以直接调用其方法
	VisitCPU(c *cpu)
	VisitDisk(d *disk)
	VisitChassis(c *chassis)
	VisitBus(b *bus)
}

// cpu 此处可以使用接口进行抽象
type cpu struct {
}

// Accept 每一个Accept方法即代表可以接收访问者对象对自身进行访问 并在其内部调用对自己的访问方法
func (c *cpu) Accept(visitor EquipmentVisitor) {
	visitor.VisitCPU(c)
}

// Core 自身方法
func (c *cpu) Core() {
	fmt.Println("Cpu Core Number is 128.")
}

type disk struct {
}

func (d *disk) Accept(visitor EquipmentVisitor) {
	visitor.VisitDisk(d)
}

func (d *disk) Capacity() {
	fmt.Println("Dick Storage Capacity is: 128T")
}

type chassis struct {
}

func (c *chassis) Accept(visitor EquipmentVisitor) {
	visitor.VisitChassis(c)
}

func (c *chassis) Info() {
	fmt.Println("Turing Gen 16 Architecture")
}

type bus struct {
}

func (b *bus) Accept(visitor EquipmentVisitor) {
	visitor.VisitBus(b)
}

func (b *bus) BandWidth() {
	fmt.Println("1T bps")
}

// infoVisitor 信息打印的访问者
type infoVisitor struct {
}

func (iv *infoVisitor) VisitCPU(c *cpu) {
	c.Core()
}

func (iv *infoVisitor) VisitDisk(d *disk) {
	d.Capacity()
}

func (iv *infoVisitor) VisitChassis(c *chassis) {
	c.Info()
}

func (iv *infoVisitor) VisitBus(b *bus) {
	b.BandWidth()
}

type computer struct {
	c  *cpu
	d  *disk
	ch *chassis
	b  *bus
}

func NewComputer() computer {
	return computer{
		&cpu{},
		&disk{},
		&chassis{},
		&bus{},
	}
}

func NewVisitor() EquipmentVisitor {
	return &infoVisitor{}
}
