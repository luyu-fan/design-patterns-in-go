package main

import __builder "github.com/luyu-fan/design-patterns-in-go/3-builder"

func main() {

	textReader := __builder.NewTextReader("Hahahah, github.com/luyu-fan/design-patterns-in-go/3-builder")
	converter := __builder.NewTexConverter()
	textReader.ParseDocument(converter)

	return
}
