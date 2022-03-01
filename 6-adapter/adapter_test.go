package __adapter

import "testing"

func TestClassAdapter(t *testing.T) {

	whiteCat := Cat{}
	kittyDog := KittyDog{&whiteCat}
	// kittyDog因为内嵌了whiteCat 实际上是可以直接调用方法的
	kittyDog.SayHello()
	kittyDog.Work()
	kittyDog.WatchDoor()

}

func TestObjectAdapter(t *testing.T) {

	whiteCat := Cat{}
	smallDog := SmallDog{&whiteCat}
	// 此时smallDog不会直接调用whiteCat的方法
	smallDog.SayHello()
	smallDog.Work()
	smallDog.WatchDoor()

}
