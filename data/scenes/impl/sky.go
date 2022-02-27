package impl

import (
	"fmt"
	_interface "github.com/luyu-fan/design-patterns-in-go/data/scenes/interface"
)

// staticSky 静态天空盒
type staticSky struct {
	effects []*_interface.Effect
	weather string
}

// DisplaySky staticSky展示天空
func (s *staticSky) DisplaySky() {
	fmt.Println("It a Static Skybox, weather is", s.weather)
}

// Info 信息打印
func (s *staticSky) Info() {
	fmt.Println("I am a staticSkyBox.")
}

// dynamicSky 动态天空盒
type dynamicSky struct {
	effects  []_interface.Effect
	skyParts []*staticSky
}

// DisplaySky dynamicSky展示动态天空
func (d *dynamicSky) DisplaySky() {
	fmt.Println("It a Dynamic Skybox.")
	for i := range d.skyParts {
		fmt.Println("Part ", i, "weather is ", d.skyParts[i].weather)
	}
}

// Info 信息打印
func (d *dynamicSky) Info() {
	fmt.Println("I am a dynamicSkyBox.")
}

// NewStaticSkyBox 静态天空盒的构造函数
func NewStaticSkyBox(w string) _interface.SkyBox {
	return &staticSky{
		weather: w,
	}
}

// NewDynamicSkyBox 动态天空盒的构造函数
func NewDynamicSkyBox(ws ...string) _interface.SkyBox {
	dsky := dynamicSky{}
	for i := range ws {
		dsky.skyParts = append(dsky.skyParts, &staticSky{weather: ws[i]})
	}
	return &dsky
}
