package impl

import (
	"fmt"
	_interface "github.com/luyu-fan/design-patterns-in-go/data/scenes/interface"
)

// GameNPC 游戏中的各类NPC
type gameNPC struct {
	desc       string
	hp         float64
	defend     float64
	attack     float64
	decorating []interface{}
}

// Info 普通对象接口
func (g *gameNPC) Info() {
	fmt.Println("I am a gameNPC:", g.desc)
	fmt.Println("decorating:", g.decorating)
}

// Clone 克隆接口
func (g *gameNPC) Clone() _interface.GameObject {
	// Go中深拷贝
	tmp := *g
	cloneNpc := &tmp
	return cloneNpc
}

// Initialize 值初始化
func (g *gameNPC) Initialize(v ...interface{}) {
	// 根据新的值进行初始化
	for _, v := range v {
		fmt.Println("NPC -> Initialize Item Value:", v)
		// 注意观察深拷贝以及切片增长时的引用问题
		g.decorating = append(g.decorating, v)
	}
}

// weapon 游戏中的武器
type weapon struct {
	desc   string
	attack float64
	weight float64
}

// Info 普通对象接口
func (w *weapon) Info() {
	fmt.Println("I am a gameWeapon:", w.desc)
}

// Clone 克隆
func (w *weapon) Clone() _interface.GameObject {
	// Go中深拷贝
	tmp := *w
	cloneWeapon := &tmp
	return cloneWeapon
}

// Initialize 值初始化
func (w *weapon) Initialize(v ...interface{}) {
	// 根据新的值进行初始化
	for _, v := range v {
		fmt.Println("Weapon -> Initialize Item Value:", v)
	}
}

// 原型注册表
var prototypeMap map[string]_interface.GameObject

// RegisterPrototypes 注册原型
func RegisterPrototypes() {
	prototypeMap = make(map[string]_interface.GameObject)
	prototypeMap["npc"] = &gameNPC{
		desc:       "game npc",
		hp:         100,
		defend:     20,
		attack:     10,
		decorating: []interface{}{1, 2, 3},
	}
	prototypeMap["weapon"] = &weapon{
		desc:   "game weapon",
		attack: 100,
		weight: 2.0,
	}
}

// GetPrototype 获取用于克隆的原型对象
func GetPrototype(kindStr string) _interface.GameObject {
	object, ok := prototypeMap[kindStr]
	if ok {
		return object
	}
	return nil
}

func init() {
	// 当原型过多的情况下可以考虑懒加载
	RegisterPrototypes()
}
