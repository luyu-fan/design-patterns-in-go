package impl

import (
	"fmt"
	_interface "github.com/luyu-fan/design-patterns-in-go/data/scenes/interface"
)

// baseWall 基础
type baseWall struct {
	materialType uint8
	hp           int32
}

// brickWall 砖墙
type brickWall struct {
	baseWall
}

// BuildWall 建砖墙
func (b *brickWall) BuildWall() {
	fmt.Println("Building a brickWall!")
}

// Info 信息打印
func (b *brickWall) Info() {
	fmt.Println("I am a brickWall")
}

// woodenWall 横木围墙
type woodenWall struct {
	baseWall
}

// BuildWall 建木制墙
func (w *woodenWall) BuildWall() {
	fmt.Println("Building a woodenWall!")
}

// Info 信息打印
func (w *woodenWall) Info() {
	fmt.Println("I am a woodenWall")
}

// NewBrickWall 砖墙构造函数
func NewBrickWall(hp int32) _interface.Wall {
	return &brickWall{
		baseWall{
			materialType: 1,
			hp:           hp,
		},
	}
}

// NewWoodenWall 横木墙构造函数
func NewWoodenWall(hp int32) _interface.Wall {
	return &woodenWall{
		baseWall{
			materialType: 2,
			hp:           hp,
		},
	}
}
