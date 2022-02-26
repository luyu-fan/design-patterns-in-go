package _interface

// Goods 通用商品接口
type Goods interface {
	// Info 信息展示方法
	Info()
	// Price 获取价格属性
	Price() float64
	// SetPrice 设置价格
	SetPrice(float64)
}

// Buyer 买家通用接口
type Buyer interface {
	Buy(g Goods)
}

// Seller 卖家通用接口
type Seller interface {
	Sell(g Goods)
	GetOne() Goods
	Add(g Goods)
}
