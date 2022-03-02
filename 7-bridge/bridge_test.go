package __bridge

import "testing"

func TestBridgeBasicWindow(t *testing.T) {

	// 可以使用工厂模式以隐藏搭配细节
	window := NewWindow("Basic")
	window.RenderAll()

}

func TestBridgeTransparentWindow(t *testing.T) {

	// 可以使用工厂模式以隐藏搭配细节
	window := NewWindow("Transparent")
	window.RenderAll()

}
