package __abstractfactory

import "testing"

func TestNewSimpleGameSceneFactory(t *testing.T) {

	// 客户在此调用抽象工厂创建具体实例 而细节已经被隐藏掉
	factory := NewSimpleGameSceneFactory()

	room := factory.BuildRoom()
	room.Info()

	skybox := factory.BuildSkyBox()
	skybox.DisplaySky()
	skybox.Info()

}

func TestNewWildSceneFactory(t *testing.T) {

	// 客户在此调用抽象工厂创建具体实例 而细节已经被隐藏掉
	factory := NewWildSceneFactory()

	room := factory.BuildRoom()
	room.Info()

	skybox := factory.BuildSkyBox()
	skybox.DisplaySky()
	skybox.Info()
}
