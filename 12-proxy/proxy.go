package _2_proxy

import "fmt"

// Graphics 图形接口
type Graphics interface {
	Draw()
	GetExtent() interface{}
	HandleMouse(event interface{})
	Load(in interface{})
	Store(out interface{})
}

type Image struct {
	// ... 一些私有状态
}

func (i *Image) Draw() {
	fmt.Println("Drawing")
}

func (i *Image) GetExtent() interface{} {
	// ... 返回图像的真实尺寸
	fmt.Println("Getting Image Extent")
	return 0
}

func (i *Image) HandleMouse(e interface{}) {
	// ... 处理鼠标事件
	fmt.Println("Handling Mouse Event", e)
}

func (i *Image) Load(in interface{}) {
	// ... 可以是自身的一些实现
}

func (i *Image) Store(out interface{}) {
	// ... 可以是自身的一些实现
}

type ImageProxy struct {
	// 维持着对真实对象的缓存或者是引用的私有状态
	image     *Image
	extent    interface{}
	imageName string
}

func (ip *ImageProxy) Draw() {
	fmt.Println("ImageProxy Draws")
	ip.getImage().Draw()
}

func (ip *ImageProxy) GetExtent() interface{} {
	fmt.Println("ImageProxy Gets Event")
	if ip.extent == nil {
		ip.extent = ip.getImage().GetExtent
	}
	return ip.extent
}

func (ip *ImageProxy) HandleMouse(e interface{}) {
	// ... 处理鼠标事件
	fmt.Println("ImageProxy Handles Mouse Event", e)
	ip.getImage().HandleMouse(e)
}

func (ip *ImageProxy) Load(in interface{}) {
	// ... 从数据流中加载图像 in应该是一个输入流
	fmt.Println("ImageProxy Loads From Stream", in)
}

func (ip *ImageProxy) Store(out interface{}) {
	// ... 将图像写入到输出流中
	fmt.Println("ImageProxy Stores To Stream", out)
}

// getImage 代理用于获取真实图像的私有方法
func (ip *ImageProxy) getImage() *Image {
	if ip.image == nil {
		// 加载真实图像
		return &Image{}
	}
	return ip.image
}

// NewImageProxy 获取图像对象
func NewImageProxy(name string) Graphics {
	return &ImageProxy{imageName: name}
}
