package _5_interpreter

// BooleanExp 定义布尔表达式的解释器
type BooleanExp interface {
	// Evaluate 计算当前结点的值
	Evaluate(ctx Context) bool
	// Replace 替换当前结点的值(用当前结点替换目标结点)
	Replace(str string, exp BooleanExp) BooleanExp
	// Copy 拷贝当前结点
	Copy() BooleanExp
}

type Context struct {
}

// Lookup 在全局结果中查找目标结点的值
func (c *Context) Lookup(str string) bool {
	// ... 根据名称查找值
	return false
}

// Assign 将某个结点的值赋为给定值
func (c *Context) Assign(exp BooleanExp, b bool) {

}

// VariableExp 变量结点
type VariableExp struct {
	name string
}

func (ve *VariableExp) Evaluate(ctx Context) bool {
	return ctx.Lookup(ve.name)
}

func (ve *VariableExp) Replace(name string, exp BooleanExp) BooleanExp {
	if name == ve.name {
		return exp.Copy()
	}
	return ve.Copy()
}

func (ve *VariableExp) Copy() BooleanExp {
	return &VariableExp{
		name: ve.name,
	}
}

// AndExp 与操作结点
type AndExp struct {
	// 定义两个操作数
	opExp1 BooleanExp
	opExp2 BooleanExp
}

func (ae *AndExp) Evaluate(ctx Context) bool {
	return ae.opExp1.Evaluate(ctx) && ae.opExp2.Evaluate(ctx)
}

func (ae *AndExp) Replace(name string, exp BooleanExp) BooleanExp {
	return &AndExp{
		ae.opExp1.Replace(name, exp),
		ae.opExp2.Replace(name, exp),
	}
}

func (ae *AndExp) Copy() BooleanExp {
	return &AndExp{
		ae.opExp1.Copy(),
		ae.opExp2.Copy(),
	}
}

// TODO 待完成后续部分
// ...
