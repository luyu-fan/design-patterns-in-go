/**
简单接口的类实现
*/

package impl

import (
	"fmt"
	"github.com/luyu-fan/design-patterns-in-go/data/goods/interface"
)

// **********************************************************************************************
// Concrete Goods
// **********************************************************************************************

// Apple 具体商品Apple
type Apple struct {
	name    string
	version string
	price   float64
	addr    string
}

// Info 实现Goods接口中的Info方法(方法名称、参数、返回值的一致性，已经方法绑定时的区别)
func (a *Apple) Info() {
	fmt.Printf("Name:%s, Version:%s, Price:%f, Addr:%s\n", a.name, a.version, a.price, a.addr)
}

// Price 获取价格
func (a *Apple) Price() float64 {
	return a.price
}

// SetPrice 设置价格
func (a *Apple) SetPrice(p float64) {
	a.price = p
}

// Nokia 具体商品Nokia
type Nokia struct {
	name      string
	version   string
	price     float64
	addr      string
	character string
}

// Info 实现Goods接口中的Info方法(方法名称、参数、返回值的一致性，已经方法绑定时的区别)
func (n *Nokia) Info() {
	fmt.Printf("Name:%s, Version:%s, Price:%f, Addr:%s, Note: Nokia has been bought by Microsoft\n", n.name, n.version, n.price, n.addr)
	fmt.Println("Character:", n.character)
}

// Price 获取价格
func (n *Nokia) Price() float64 {
	return n.price
}

// SetPrice 设置价格
func (n *Nokia) SetPrice(p float64) {
	n.price = p
}

// Deal 交易
func (n *Nokia) Deal(b _interface.Buyer, s _interface.Seller) {
	s.Sell(n)
	b.Buy(n)
	fmt.Println("Nokia Deal finished!")
}

// **********************************************************************************************
// Concrete Buyers
// **********************************************************************************************

// ChineseCustomer 中国顾客
type ChineseCustomer struct {
	money float64
	rate  float64
	name  string
	goods []_interface.Goods
}

// Buy 实现Buyer接口
func (c *ChineseCustomer) Buy(g _interface.Goods) {
	fmt.Println("==============================================================================")
	fmt.Printf("买家%s正在买入:\n", c.name)
	g.Info()
	cost := g.Price() * c.rate
	fmt.Println("花费:", cost)
	c.money -= cost
	fmt.Println("剩余:", c.money)
	c.goods = append(c.goods, g)
	fmt.Println("买入完毕!")
	fmt.Println("==============================================================================")
}

// UsaCustomer 美国顾客
type UsaCustomer struct {
	money float64
	name  string
	goods []_interface.Goods
}

// Buy 实现Buyer接口
func (u *UsaCustomer) Buy(g _interface.Goods) {
	fmt.Println("==============================================================================")
	fmt.Printf("Customer %s is purchasing:\n", u.name)
	g.Info()
	u.money -= g.Price()
	u.goods = append(u.goods, g)
	fmt.Println("Buying finished!")
	fmt.Println("==============================================================================")
}

// **********************************************************************************************
// Concrete Sellers
// **********************************************************************************************

type BeijingDigitStore struct {
	name     string
	goods    []_interface.Goods
	addr     string
	turnover float64
	tariff   float64
}

type NewYorkDigitStore struct {
	name     string
	goods    []_interface.Goods
	addr     string
	turnover float64
}

// Sell 实现Seller接口
func (b *BeijingDigitStore) Sell(g _interface.Goods) {
	b.goods = b.goods[1:]
	g.SetPrice(g.Price() + b.tariff)
	b.turnover += g.Price()
}

// GetOne 实现Seller接口
func (b *BeijingDigitStore) GetOne() _interface.Goods {
	if len(b.goods) > 0 {
		return b.goods[0]
	}
	return nil
}

// Add 实现Seller接口
func (b *BeijingDigitStore) Add(g _interface.Goods) {
	if b.goods == nil {
		b.goods = make([]_interface.Goods, 0)
	}
	b.goods = append(b.goods, g)
}

// Sell 实现Seller接口
func (n *NewYorkDigitStore) Sell(g _interface.Goods) {
	n.goods = n.goods[1:]
	n.turnover += g.Price()
}

// GetOne 实现Seller接口
func (n *NewYorkDigitStore) GetOne() _interface.Goods {
	if len(n.goods) > 0 {
		return n.goods[0]
	}
	return nil
}

// Add 实现Seller接口
func (n *NewYorkDigitStore) Add(g _interface.Goods) {
	if n.goods == nil {
		n.goods = make([]_interface.Goods, 0)
	}
	n.goods = append(n.goods, g)
}

// **********************************************************************************************
// Constructors
// **********************************************************************************************

// NewApple Apple构造函数
func NewApple(v string, p float64) Apple {
	return Apple{
		name:    "Apple",
		version: v,
		price:   p,
		addr:    "Street 13, Saint Los Angeles",
	}
}

// NewNokia Apple构造函数
func NewNokia(v string, p float64) Nokia {
	return Nokia{
		name:      "Nokia",
		version:   v,
		price:     p,
		addr:      "Street 8, Saint Los Angeles",
		character: "Microsoft has bought Nokia Inc.",
	}
}

// NewChineseCustomer ChineseCustomer构造函数
func NewChineseCustomer(n string, m float64) ChineseCustomer {
	return ChineseCustomer{
		name:  n,
		rate:  6.67,
		money: m,
	}
}

// NewUsaCustomer UsaCustomer构造函数
func NewUsaCustomer(n string, m float64) UsaCustomer {
	return UsaCustomer{
		name:  n,
		money: m,
	}
}

// NewBeijingDigitStore BeijingDigitStore 构造函数
func NewBeijingDigitStore(n, addr string, t float64) BeijingDigitStore {
	return BeijingDigitStore{
		name:   n,
		addr:   addr,
		tariff: t,
	}
}

// NewNewYorkDigitStore NewYorkDigitStore 构造函数
func NewNewYorkDigitStore(n, addr string) NewYorkDigitStore {
	return NewYorkDigitStore{
		name: n,
		addr: addr,
	}
}
