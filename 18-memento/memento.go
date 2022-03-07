package _8_memento

import (
	"fmt"
	"math/rand"
)

type Pos [2]int

// Car 原发器 内部完成对象功能的实现
type Car struct {
	pos Pos
}

func (c *Car) Move() {
	c.pos[0] += rand.Intn(10)
	c.pos[1] += rand.Intn(10)
	fmt.Println("Move Car to:", c.pos)
}

func (c *Car) CreatePosMemento() PosMemento {
	return PosMemento{
		c.pos,
	}
}

func (c *Car) SetMemento(m PosMemento) {
	c.pos = m.GetState()
	fmt.Println("Reset pos: ", c.pos)
}

// MoveCommand 可以使用更抽象的接口 在这里是
type MoveCommand struct {
	target Car
	memos  []PosMemento
}

// Execute 执行操作
func (mc *MoveCommand) Execute() {
	memo := mc.target.CreatePosMemento()
	mc.memos = append(mc.memos, memo)
	mc.target.Move()
}

// UnExecute 撤销操作
func (mc *MoveCommand) UnExecute() {
	if len(mc.memos) == 0 {
		fmt.Println("Nothing to do")
		return
	}
	lastMemo := mc.memos[len(mc.memos)-1]
	mc.memos = mc.memos[0 : len(mc.memos)-1]
	mc.target.SetMemento(lastMemo)
}

// PosMemento 一个简单的备忘录
type PosMemento struct {
	// 保存内部数据
	pos Pos
	// ... 其它状态数据
}

func (pm *PosMemento) GetState() Pos {
	return pm.pos
}

func (pm *PosMemento) SetState(pos Pos) {
	pm.pos = pos
}
