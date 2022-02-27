package impl

import (
	"fmt"
	_interface "github.com/luyu-fan/design-patterns-in-go/data/scenes/interface"
)

type baseDoor struct {
	aRoom _interface.Room
	bRoom _interface.Room
}

type steelDoor struct {
	baseDoor
	effects  []_interface.Effect
	property string
}

// BuildDoor 铁门的建造过程
func (s *steelDoor) BuildDoor() {
	fmt.Println("Building a steelDoor. Do Some Visual effects")
}

// Connect 铁门的安装过程
func (s *steelDoor) Connect(r1, r2 _interface.Room) {
	s.aRoom = r1
	s.bRoom = r2
	fmt.Println("steelDoor is Connecting Rooms. Do Some Visual effects")
}

// Info 信息打印
func (s *steelDoor) Info() {
	fmt.Println("I am a SteelDoor, my property is ", s.property)
}

// woodenDoor 木质门
type woodenDoor struct {
	baseDoor
	property string
}

// BuildDoor 木质门的建造过程
func (w *woodenDoor) BuildDoor() {
	fmt.Println("Building a woodenDoor.")
}

// Connect 木质门的安装过程
func (w *woodenDoor) Connect(r1, r2 _interface.Room) {
	w.aRoom = r1
	w.bRoom = r2
	fmt.Println("woodenDoor is Connecting Rooms.")
}

// Info 信息打印
func (w *woodenDoor) Info() {
	fmt.Println("I am a WoodenDoor, my property is ", w.property)
}

// glassDoor 玻璃门
type glassDoor struct {
	baseDoor
	effects  []_interface.Effect
	property string
	opacity  float64
}

// BuildDoor 玻璃门建造过程
func (g *glassDoor) BuildDoor() {
	fmt.Println("Building a glassDoor.")
	fmt.Println("Be Careful!")
}

// Connect 玻璃门安装过程
func (g *glassDoor) Connect(r1, r2 _interface.Room) {
	g.aRoom = r1
	g.bRoom = r2
	fmt.Println("glassDoor is Connecting Rooms.")
	fmt.Println("Be Careful!")
}

// Info 信息打印
func (g *glassDoor) Info() {
	fmt.Println("I am a GlassDoor, my property is ", g.property)
}

// NewWoodenDoor woodenDoor 构造函数
func NewWoodenDoor() _interface.Door {
	return &woodenDoor{
		property: "flammable",
	}
}

// NewSteelDoor steelDoor构造函数
func NewSteelDoor() _interface.Door {
	return &steelDoor{
		property: "bulky and rusty",
	}
}

// NewGlassDoor glassDoor构造函数
func NewGlassDoor() _interface.Door {
	return &glassDoor{
		property: "fragile, transparent",
		opacity:  0.5,
	}
}
