package inherit

import "testing"

func TestBinding(t *testing.T) {

	// 结构体变量
	shape1 := Shape{}
	shape1.Show()
	shape1.StructShow()

	t.Log("=======================================")

	shape1.ModifyB(456456)
	shape1.Show()
	shape1.StructShow()

	t.Log("=======================================")
	// 结构体变量调用指针接收者依然能够修改其内部值
	shape1.ModifyA(789789)
	shape1.Show()
	shape1.StructShow()

	t.Log("=======================================")

	// 指针变量
	shape2 := &Shape{}
	shape2.Show()
	shape2.StructShow()

	t.Log("=======================================")
	// 指针对象调用值绑定不会修改内部值
	shape2.ModifyB(456456)
	shape2.Show()
	shape2.StructShow()

	t.Log("=======================================")
	// 指针对象调用指针绑定会修改内部对象
	shape2.ModifyA(789789)
	shape2.Show()
	shape2.StructShow()

	// 总结:
	// 1. go 1.16对于结构体接收者/指针接收者 无论是结构体变量还是指针变量都可以调用
	// 2. 无论是值绑定还是指针调用 只要绑定的为指针方法即可修改内部数据

}

// TestRectangle 测试内部是内嵌结构体的情况下有什么影响
func TestRectangle(t *testing.T) {

	baseShape := Shape{serialNumber: 0}

	rigidShape := RigidShape{
		Shape: &baseShape,
		color: 123,
	}

	rectangle := Rectangle{
		RigidShape: &rigidShape,
		width:      12,
		height:     32,
	}

	// ======================================================================================
	// !!! 结论: 方法调用依然遵循方法绑定的那一套规则，在结构体继承的情况下数值修改内外相互隔绝 而指针继承会互相影响
	t.Log("==================== Modifying from top to down(Outer) ====================")
	baseShape.serialNumber = 1234
	rectangle.Show()

	//rigidShape.color = 4321
	//rectangle.Show()
	rectangle.GetColor()
	rigidShape.color = 6666
	rectangle.GetColor()
	rigidShape.GetColor()

	t.Log("==================== Modifying from top to down(Inner) ====================")
	rectangle.serialNumber = 667667
	rectangle.Show()
	rectangle.color = 768678
	rectangle.GetColor()
	rigidShape.GetColor()
}
