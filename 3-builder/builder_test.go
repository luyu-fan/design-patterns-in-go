package __builder

import (
	"reflect"
	"testing"
)

func TestNewAsciiConverter(t *testing.T) {
	converter := NewAsciiConverter()
	t.Log(reflect.TypeOf(converter))
	t.Log(reflect.ValueOf(converter))
}

func BenchmarkNewAsciiConverter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		converter := NewAsciiConverter()
		b.Log(reflect.TypeOf(converter))
		b.Log(reflect.ValueOf(converter))
	}
}

func TestNewTexConverter(t *testing.T) {
	converter := NewTexConverter()
	t.Log(reflect.TypeOf(converter))
	t.Log(reflect.ValueOf(converter))
}

func BenchmarkNewTexConverter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		converter := NewAsciiConverter()
		b.Log(reflect.TypeOf(converter))
		b.Log(reflect.ValueOf(converter))
	}
}

func TestTextReaderParseDocument(t *testing.T) {
	textReader := NewTextReader("Builder将功能以及作用相似的一族对象的构建过程或某个功能算法过程抽象出来，划分若干个明显的步骤，每种对象对应的类按照规定的步骤实现以达到不同的表示。")
	converter := NewAsciiConverter()

	// 客户仅将converter作为参数传入
	textReader.ParseDocument(converter)
}

func BenchmarkTextReaderParseDocument(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textReader := NewTextReader("Builder将功能以及作用相似的一族对象的构建过程或某个功能算法过程抽象出来，划分若干个明显的步骤，每种对象对应的类按照规定的步骤实现以达到不同的表示。")
		converter := NewAsciiConverter()

		// 客户仅将converter作为参数传入
		textReader.ParseDocument(converter)
	}
}
