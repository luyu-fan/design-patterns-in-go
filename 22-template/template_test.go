package _2_template

import "testing"

func TestTemplate(t *testing.T) {
	concrete := NewClass()

	// 使用具体类调用 (并不能达到目的)
	concrete.TemplateMethod()
}
