package __adapter

import "fmt"

type DogAnimal interface {
	WatchDoor()
	Work()
	SayHello()
}

type CatAnimal interface {
	Play()
	CatchMouse()
	Flighty()
}

type Cat struct {
}

func (c *Cat) Play() {
	fmt.Println("Playing!")
}

func (c *Cat) CatchMouse() {
	fmt.Println("Catching Mouse!")
}

func (c *Cat) Flighty() {
	fmt.Println("Flighty!")
}

// KittyDog 类适配器 适配Cat这一个类 对于Go而言 类适配器不如对象适配器更强大
type KittyDog struct {
	*Cat
}

func (cd *KittyDog) WatchDoor() {
	fmt.Println("I am a kitty dog! I am watching door!")
	cd.Play()
}

func (cd *KittyDog) Work() {
	fmt.Println("I am a kitty dog! I am working")
	cd.CatchMouse()
}

func (cd *KittyDog) SayHello() {
	fmt.Println("I am a kitty dog! I am saying hello")
	cd.Flighty()
}

// SmallDog 对象适配器 适配CatAnimal类型
type SmallDog struct {
	innerCatAnimal CatAnimal
}

func (sd *SmallDog) WatchDoor() {
	fmt.Println("I am a small dog, I am watching door!")
	sd.innerCatAnimal.Play()
}

func (sd *SmallDog) Work() {
	fmt.Println("I am a small dog, I am working!")
	sd.innerCatAnimal.CatchMouse()
}

func (sd *SmallDog) SayHello() {
	fmt.Println("I am a small dog, I am saying hello!")
	sd.innerCatAnimal.Flighty()
}
