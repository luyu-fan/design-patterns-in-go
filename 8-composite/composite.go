package __composite

import "fmt"

// View 统一结构
type View interface {
	Draw()
	Add(v View)
	Remove(v View)
	IsComponent() bool
}

// panel 容器
type panel struct {
	components []View
	// 此处可以维护父指针
}

func (p *panel) Draw() {
	fmt.Println("Panel Drawing!")
	for i := range p.components {
		p.components[i].Draw()
	}
}

func (p *panel) Add(v View) {
	fmt.Println("Panel Adds Node:", v)
	p.components = append(p.components, v)
}

func (p *panel) Remove(v View) {
	// 根据v查找具体的节点并删除
	// 这里可以考虑使用LRU等缓存结构
	fmt.Println("Panel Removes Node:", v)
}

func (p *panel) IsComponent() bool {
	return true
}

// button 叶子
type button struct {
}

func (b *button) Draw() {
	fmt.Println("Button Drawing!")
}

func (b *button) Add(v View) {
	if !b.IsComponent() {
		panic("Button can not add!")
	}
}

func (b *button) Remove(v View) {
	if !b.IsComponent() {
		panic("Button can not remove!")
	}
}

func (b *button) IsComponent() bool {
	return false
}

// label 叶子
type label struct {
}

func (l *label) Draw() {
	fmt.Println("Label Drawing!")
}

func (l *label) Add(v View) {
	if !l.IsComponent() {
		panic("Label can not add!")
	}
}

func (l *label) Remove(v View) {
	if !l.IsComponent() {
		panic("Label can not Remove!")
	}
}

func (l *label) IsComponent() bool {
	return false
}
