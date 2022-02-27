package __abstractfactory

import _interface "github.com/luyu-fan/design-patterns-in-go/data/scenes/interface"

// GameScene 游戏场景抽象工厂
type GameScene interface {
	BuildRoom() _interface.Room
	BuildSkyBox() _interface.SkyBox
}
