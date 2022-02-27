package main

import (
	"github.com/luyu-fan/design-patterns-in-go/data/goods/impl"
)

func main() {

	taiguliStore := impl.NewBeijingDigitStore("太古里专卖店", "北京市三里屯太古里", 120)
	xiaoming := impl.NewChineseCustomer("小明", 12000.0)
	lisa := impl.NewUsaCustomer("Lisa", 1000.0)

	applePhone1 := impl.NewApple("Meg, 14.0174", 894.6)
	nokiaPhone1 := impl.NewNokia("Ming, 6.8", 322.0)

	taiguliStore.Add(&applePhone1)
	taiguliStore.Add(&nokiaPhone1)

	phone := taiguliStore.GetOne()
	taiguliStore.Sell(phone)
	xiaoming.Buy(phone)

	phone = taiguliStore.GetOne()
	taiguliStore.Sell(phone)
	lisa.Buy(phone)

	return
}
