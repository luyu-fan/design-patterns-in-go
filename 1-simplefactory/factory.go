package simplefactory

import (
	"github.com/luyu-fan/design-patterns-in-go/data/goods/impl"
	_interface "github.com/luyu-fan/design-patterns-in-go/data/goods/interface"
	"sync"
)

// ap 全局工厂变量
var ap *ApplePhoneFactory

// once 全局锁
var once = &sync.Once{}

// ApplePhoneFactory 简单工厂
type ApplePhoneFactory struct {
}

// CreatePhone 主要工厂方法 隐藏Phone实现的细节
func (a *ApplePhoneFactory) CreatePhone(n string, pr float64) _interface.Goods {
	phone := impl.NewApple(n, pr)
	return &phone
}

// NewApplePhoneFactory ApplePhoneFactory构造方法 一般为全局公有 采用单例模式
func NewApplePhoneFactory() *ApplePhoneFactory {
	// 采用lazy模式创建全局工厂变量
	once.Do(func() {
		ap = &ApplePhoneFactory{}
	})
	return ap
}
