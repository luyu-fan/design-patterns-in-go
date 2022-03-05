package _2_proxy

import "testing"

func TestNewImageProxy(t *testing.T) {

	img := NewImageProxy("photo.jpg")

	img.GetExtent()

	img.Draw()

	img.HandleMouse("Click")

	img.Store(nil)
	img.Load(nil)
}
