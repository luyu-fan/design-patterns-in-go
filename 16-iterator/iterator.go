package _6_iterator

// List 列表接口
type List interface {
	Count() int
	Get(index int) interface{}
	Append(v interface{})
	Pop(index int) interface{}
	// 其它的一些和列表相关的实现
}

// Iterator 迭代器接口(外部迭代器)
type Iterator interface {
	// First 移动到开始位置
	First() interface{}
	// Back 结尾位置
	Back() interface{}
	// Next 下一个元素
	Next() interface{}
	// Length 获取长度
	Length() int
	// IsDone 是否访问完毕
	IsDone() bool
}

type simpleList struct {
	innerArray []interface{}
}

func (sl *simpleList) Count() int {
	return len(sl.innerArray)
}

func (sl *simpleList) Get(index int) interface{} {
	return sl.innerArray[index]
}

func (sl *simpleList) Append(v interface{}) {
	sl.innerArray = append(sl.innerArray, v)
}

func (sl *simpleList) Pop(index int) interface{} {
	v := sl.innerArray[index]
	sl.innerArray = append(sl.innerArray[:index], sl.innerArray[index+1:])
	return v
}

type simpleIterator struct {
	current int
	list    List
}

func (si *simpleIterator) First() interface{} {
	si.current = 0
	return si.list.Get(si.current)
}

func (si *simpleIterator) Back() interface{} {
	si.current = si.Length() - 1
	return si.list.Get(si.current)
}

func (si *simpleIterator) Next() interface{} {
	v := si.list.Get(si.current)
	si.current += 1
	return v
}

func (si *simpleIterator) Length() int {
	return si.list.Count()
}

func (si *simpleIterator) IsDone() bool {
	if si.Length() == si.current {
		return true
	}
	return false
}

func NewSimpleList(v ...interface{}) List {
	return &simpleList{
		v,
	}
}

func NewSimpleIterator(l List) Iterator {
	return &simpleIterator{list: l}
}

// 与之对应的是内部遍历器，即首先引用迭代器对象 然后在遍历时先对每一个对象进行处理，以子类实现的形式完成自定义功能。
