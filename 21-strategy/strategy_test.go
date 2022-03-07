package _1_strategy

import "testing"

func TestStrategy(t *testing.T) {

	dc := NewDiscountCenter()

	myCarts := NewShoppingCart()

	xiaoming := Customer{carts: myCarts}

	// 使用不同的算法
	xiaoming.Pay(dc.GetAlg("basic"))
	xiaoming.Pay(dc.GetAlg("christmas"))

}
