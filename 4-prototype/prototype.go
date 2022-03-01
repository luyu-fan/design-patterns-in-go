package __prototype

import "github.com/luyu-fan/design-patterns-in-go/data/scenes/impl"

type Client struct {
}

// GetPrototype 获取原型
func (c *Client) GetPrototype() {
	object := impl.GetPrototype("npc")
	object.Info()
	object = impl.GetPrototype("weapon")
	object.Info()
}

// GetPrototypeAndCopy 克隆原型
func (c *Client) GetPrototypeAndCopy() {
	proto := impl.GetPrototype("npc")
	proto.Info()
	newObject := proto.Clone()
	newObject.Initialize("4")
	// 观察Go中的拷贝操作对Slice的影响
	proto.Info()
	newObject.Info()
}
