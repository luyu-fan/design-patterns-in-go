package __abstractfactory

import (
	"github.com/luyu-fan/design-patterns-in-go/data/scenes/impl"
	_interface "github.com/luyu-fan/design-patterns-in-go/data/scenes/interface"
)

// simpleGameSceneFactory 构建简单房间的简单工厂类
type simpleGameSceneFactory struct {
}

// BuildRoom 构建房间
func (sg *simpleGameSceneFactory) BuildRoom() _interface.Room {

	room := impl.NewChalet()
	room.BuildRoom()

	door := impl.NewWoodenDoor()
	room.AddDoor(door)

	wall := impl.NewBrickWall(1234)
	room.AddWall(wall)

	return room

}

// BuildSkyBox 构建天空盒
func (sg *simpleGameSceneFactory) BuildSkyBox() _interface.SkyBox {
	skybox := impl.NewStaticSkyBox("Sunny")
	return skybox
}

// wildSceneFactory 构建野外场景
type wildSceneFactory struct {
}

// BuildRoom 构建房间
func (wg *wildSceneFactory) BuildRoom() _interface.Room {

	room := impl.NewChalet()
	room.BuildRoom()

	door := impl.NewWoodenDoor()
	room.AddDoor(door)

	wall := impl.NewWoodenWall(668)
	room.AddWall(wall)

	return room

}

// BuildSkyBox 构建天空盒
func (wg *wildSceneFactory) BuildSkyBox() _interface.SkyBox {
	skybox := impl.NewDynamicSkyBox("Sunny", "Cloudy", "Rainy", "Storm")
	return skybox
}

// ... plenty of concrete factories to be continued

// NewSimpleGameSceneFactory simpleGameSceneFactory简单工厂单例
func NewSimpleGameSceneFactory() GameScene {
	return sFactory
}

// NewWildSceneFactory wildSceneFactory简单工厂单例
func NewWildSceneFactory() GameScene {
	return wFactory
}

var sFactory *simpleGameSceneFactory
var wFactory *wildSceneFactory

func init() {
	sFactory = &simpleGameSceneFactory{}
	wFactory = &wildSceneFactory{}
}
