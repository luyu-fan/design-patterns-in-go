package main

import inherit "github.com/luyu-fan/design-patterns-in-go/inherit"

func main() {

	// 结构体变量
	shape1 := inherit.Shape{}
	shape1.Show()
	shape1.StructShow()

	shape1.ModifyB(456456)
	shape1.Show()
	shape1.StructShow()

	// 结构体变量调用指针接收者依然能够修改其内部值
	shape1.ModifyA(789789)
	shape1.Show()
	shape1.StructShow()

	// 指针变量
	shape2 := &inherit.Shape{}
	shape2.Show()
	shape2.StructShow()

	// 指针对象调用值绑定不会修改内部值
	shape2.ModifyB(456456)
	shape2.Show()
	shape2.StructShow()

	// 指针对象调用指针绑定会修改内部对象
	shape2.ModifyA(789789)
	shape2.Show()
	shape2.StructShow()

	return
}
