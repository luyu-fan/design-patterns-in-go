package _1_strategy

import "fmt"

type CommodityType uint8

const (
	CLOTH CommodityType = iota + 1
	ELECTRON
	DRINK
	FOOD
	BOOK
	// ... 其它种类
)

// Commodity 商品接口
type Commodity interface {
	GetName() string
	GetPrice() float64
	GetType() CommodityType
}

type milk struct {
	name  string
	price float64
	kind  CommodityType
}

func (m *milk) GetName() string {
	return m.name
}

func (m *milk) GetPrice() float64 {
	return m.price
}

func (m *milk) GetType() CommodityType {
	return m.kind
}

type tShirt struct {
	name  string
	price float64
	kind  CommodityType
}

func (t *tShirt) GetName() string {
	return t.name
}

func (t *tShirt) GetPrice() float64 {
	return t.price
}

func (t *tShirt) GetType() CommodityType {
	return t.kind
}

type xiaomi12 struct {
	name  string
	price float64
	kind  CommodityType
}

func (x *xiaomi12) GetName() string {
	return x.name
}

func (x *xiaomi12) GetPrice() float64 {
	return x.price
}

func (x *xiaomi12) GetType() CommodityType {
	return x.kind
}

// ... Others

// shoppingCart 简单购物车
type shoppingCart struct {
	things []Commodity
}

func NewShoppingCart() *shoppingCart {
	return &shoppingCart{
		[]Commodity{
			&milk{
				"Mengniu",
				12.0,
				FOOD,
			},
			&milk{
				"Yili",
				11.0,
				FOOD,
			},
			&tShirt{
				"T-shirt",
				99.0,
				CLOTH,
			},
			&xiaomi12{
				"Xiao Mi 12",
				3488.0,
				ELECTRON,
			},
		},
	}
}

// DiscountAlgorithm 折扣接口
type DiscountAlgorithm interface {
	// CalcTotalCost 计算价钱
	CalcTotalCost(carts *shoppingCart) float64
	// 可以设计其它更多的方法
}

type simpleDiscountAlg struct {
	// ... 此处可以包含其它外部必要数据
}

// CalcTotalCost 实现接口
func (sa *simpleDiscountAlg) CalcTotalCost(carts *shoppingCart) float64 {
	totalCost := 0.0
	baseBonus := 6.0
	fmt.Println("========================================================")
	for i := range carts.things {
		commodity := carts.things[i]
		fmt.Println(commodity.GetName(), commodity.GetPrice())
		totalCost += commodity.GetPrice()
	}
	if totalCost >= 90 {
		fmt.Println("Bonus: -", baseBonus)
	}
	return totalCost - baseBonus
}

type merryChristmasDiscountAlg struct {
	// ... 此处可以包含其它外部必要数据
}

func (md *merryChristmasDiscountAlg) CalcTotalCost(carts *shoppingCart) float64 {
	fmt.Println("========================================================")
	totalCost := 0.0
	for i := range carts.things {
		commodity := carts.things[i]
		fmt.Println(commodity.GetName(), commodity.GetPrice(), "Discount:", commodity.GetPrice()*0.15)
		totalCost += commodity.GetPrice() * 0.75
	}
	return totalCost
}

// discountCenter 折扣中心
type discountCenter struct {
	// algs 维护多个折扣算法
	algs map[string]DiscountAlgorithm
}

// GetAlg 获取打折算法
func (dc *discountCenter) GetAlg(name string) DiscountAlgorithm {
	return dc.algs[name]
}

// Customer 简单顾客
type Customer struct {
	carts *shoppingCart
}

// Pay 付款
func (c *Customer) Pay(alg DiscountAlgorithm) {
	// 用户可以选择合适的打折策略
	totalCost := alg.CalcTotalCost(c.carts)
	fmt.Println("It will cost: ", totalCost)
}

func NewDiscountCenter() *discountCenter {
	return &discountCenter{
		algs: map[string]DiscountAlgorithm{
			"basic":     &simpleDiscountAlg{},
			"christmas": &merryChristmasDiscountAlg{},
			// Others
		},
	}
}
