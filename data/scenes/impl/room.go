package impl

import (
	"fmt"
	_interface "github.com/luyu-fan/design-patterns-in-go/data/scenes/interface"
)

type baseRoom struct {
	doors []_interface.Door
	walls []_interface.Wall
	HP    int32
}

// chalet 木屋
type chalet struct {
	baseRoom
	effects  []_interface.Effect
	property string
}

// BuildRoom 建造木屋
func (c *chalet) BuildRoom() {
	fmt.Println("chalet! Building!")
}

// AddDoor 安装门
func (c *chalet) AddDoor(d _interface.Door) {
	c.doors = append(c.doors, d)
	fmt.Println("Bang Bang Bang! chalet is Building a Door!")
}

// AddWall 砌墙
func (c *chalet) AddWall(w _interface.Wall) {
	c.walls = append(c.walls, w)
	fmt.Println("Pa Pa Pa! chalet is Building a Wall!")
}

// Info 信息打印
func (c *chalet) Info() {
	fmt.Println("I am a Chalet. My Components are as followed:")
	for i := range c.doors {
		// 转换为GameObject对象
		c.doors[i].Info()
	}
	for i := range c.walls {
		// 转换为GameObject对象
		c.walls[i].Info()
	}
}

// apartment 公寓
type apartment struct {
	baseRoom
	effects  []_interface.Effect
	property string
}

// BuildRoom 建造公寓
func (a *apartment) BuildRoom() {
	fmt.Println("apartment! Building!")
}

// AddDoor 安装门
func (a *apartment) AddDoor(d _interface.Door) {
	a.doors = append(a.doors, d)
	fmt.Println("Bang Bang Bang! apartment is Building a Door!")
}

// AddWall 砌墙
func (a *apartment) AddWall(w _interface.Wall) {
	a.walls = append(a.walls, w)
	fmt.Println("Pa Pa Pa! apartment is Building a Wall!")
}

// Info 信息打印
func (a *apartment) Info() {
	fmt.Println("I am An apartment. My Components are as followed:")
	for i := range a.doors {
		// 转换为GameObject对象
		a.doors[i].Info()
	}
	for i := range a.walls {
		// 转换为GameObject对象
		a.walls[i].Info()
	}
}

// palace 宫殿
type palace struct {
	baseRoom
	effects  []_interface.Effect
	property string
}

// BuildRoom 建造公寓
func (p *palace) BuildRoom() {
	fmt.Println("palace! Building!")
}

// AddDoor 安装门
func (p *palace) AddDoor(d _interface.Door) {
	p.doors = append(p.doors, d)
	fmt.Println("Bang Bang Bang! palace is Building a Door!")
}

// AddWall 砌墙
func (p *palace) AddWall(w _interface.Wall) {
	p.walls = append(p.walls, w)
	fmt.Println("Pa Pa Pa! palace is Building a Wall!")
}

// Info 信息打印
func (p *palace) Info() {
	fmt.Println("I am a Palace. My Components are as followed:")
	for i := range p.doors {
		// 转换为GameObject对象
		p.doors[i].Info()
	}
	for i := range p.walls {
		// 转换为GameObject对象
		p.walls[i].Info()
	}
}

// NewChalet 木屋构造函数
func NewChalet() _interface.Room {
	return &chalet{
		property: "flammable",
	}
}

// NewApartment 公寓构造函数
func NewApartment() _interface.Room {
	return &apartment{
		property: "modern",
	}
}

// NewPalace 宫殿构造函数
func NewPalace() _interface.Room {
	return &palace{
		property: "priceless",
	}
}
