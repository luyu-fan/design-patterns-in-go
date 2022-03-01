package __singleton

import (
	"fmt"
	"sync"
)

// Factory 抽象工厂接口
type Factory interface {
	Create()
}

type factory1 struct {
	flag uint8
}

func (f1 *factory1) Create() {
	fmt.Println("Factory1 Creating...")
}

type factory2 struct {
	flag uint8
}

func (f2 *factory2) Create() {
	fmt.Println("Factory2 Creating...")
}

// 全局变量声明
var factoryMap map[string]Factory
var initOnce *sync.Once = &sync.Once{}

// Instance 在没有类提供Instance方法的情况下使用包中提供公开方法的方式代替
func Instance(singleton string) Factory {

	//// 预加载(和init配合使用)
	//return factoryMap[singleton]

	// 懒加载(和once配合使用)
	if factoryMap == nil {
		// Double-Check
		initOnce.Do(func() {
			if factoryMap == nil {
				register()
			}
		})
	}
	return factoryMap[singleton]
}

// register 注册
func register() {
	factoryMap = make(map[string]Factory)
	factoryMap["f1"] = &factory1{}
	factoryMap["f2"] = &factory2{}
}

//func init() {
//	// 预先创建各个singleton实例以供访问
//	fmt.Println("Singleton Package Init")
//	register()
//}
