package __builder

import (
	"fmt"
	"reflect"
)

// TextConverter 文档字体转换器Builder
type TextConverter interface {
	ConvertChar(char rune)
	ConvertFont(v interface{})
	ConvertParagraph(v interface{})
	// 此处可以不约定获取内容的接口方法 原因是不同的产品之间差异可能多大 强制抽象可能没有意义
}

// 产品
type asciiDocument []interface{}
type texDocument []interface{}

// asciiConverter 针对Ascii文本格式的转换
type asciiConverter struct {
	asciiDocument
}

// ConvertChar asciiConverter对单个字的转换
func (a *asciiConverter) ConvertChar(char rune) {
	a.asciiDocument = append(a.asciiDocument, char)
	fmt.Println("AsciiConverter对单个字的转换:", char)
}

// ConvertFont asciiConverter对字体的转换
func (a *asciiConverter) ConvertFont(v interface{}) {
	a.asciiDocument = append(a.asciiDocument, v)
	vType := reflect.TypeOf(v)
	vValue := reflect.ValueOf(v)
	fmt.Println("AsciiConverter对字体的转换:", vType, vValue)
}

// ConvertParagraph asciiConverter对段落的转换
func (a *asciiConverter) ConvertParagraph(v interface{}) {
	a.asciiDocument = append(a.asciiDocument, v)
	vType := reflect.TypeOf(v)
	vValue := reflect.ValueOf(v)
	fmt.Println("AsciiConverter对段落的转换:", vType, vValue)
}

// GetAsciiDocument asciiConverter转换后的文档
func (a *asciiConverter) GetAsciiDocument() asciiDocument {
	// 假设此时完成了转换 返回转换后的结果
	return a.asciiDocument
}

// texConverter 针对Tex格式的转换
type texConverter struct {
	texDoc texDocument
}

// ConvertChar texConverter对单个字的转换
func (t *texConverter) ConvertChar(char rune) {
	t.texDoc = append(t.texDoc, char)
	fmt.Println("AsciiConverter对单个字的转换:", char)
}

// ConvertFont texConverter对字体的转换
func (t *texConverter) ConvertFont(v interface{}) {
	t.texDoc = append(t.texDoc, v)
	vType := reflect.TypeOf(v)
	vValue := reflect.ValueOf(v)
	fmt.Println("TexConverter对字体的转换:", vType, vValue)
}

// ConvertParagraph texConverter对段落的转换
func (t *texConverter) ConvertParagraph(v interface{}) {
	t.texDoc = append(t.texDoc, v)
	vType := reflect.TypeOf(v)
	vValue := reflect.ValueOf(v)
	fmt.Println("TexConverter对段落的转换:", vType, vValue)
}

// GetTexDocument texConverter转换后的文档
func (t *texConverter) GetTexDocument() texDocument {
	// 假设此时完成了转换 返回转换后的结果
	return t.texDoc
}

type textReader struct {
	rawData string
}

// ParseDocument TextReader根据实际的文档转换器转换并打印
func (t *textReader) ParseDocument(converter TextConverter) {
	// 模拟对原始文档进行转换并打印
	for i, token := range []rune(t.rawData) {
		if i%3 == 0 {
			converter.ConvertChar(token)
		} else if i%2 == 0 {
			converter.ConvertFont("宋体")
		} else {
			para := "Paragraph xxx, token is:" + string(token)
			converter.ConvertParagraph(para)
		}
	}

	// 在没有对各个Builder实现约定返回接口的情况下 需要使用类型断言 从而获取具体实例
	// 但从代码角度而言 提供一个统一的返回值接口会好一些
	switch converter.(type) {
	case *asciiConverter:
		{
			asciiConvert := converter.(*asciiConverter)
			doc := asciiConvert.GetAsciiDocument()
			fmt.Println(doc)
		}
	case *texConverter:
		{
			texConvert := converter.(*texConverter)
			doc := texConvert.GetTexDocument()
			fmt.Println(doc)
		}
	}
}

// NewTextReader TextReader构造函数
func NewTextReader(text string) *textReader {
	return &textReader{text}
}

// NewAsciiConverter 构造函数
func NewAsciiConverter() TextConverter {
	return &asciiConverter{}
}

// NewTexConverter 构造函数
func NewTexConverter() TextConverter {
	return &texConverter{}
}
