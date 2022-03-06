package _6_iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {

	l := NewSimpleList(1, 2, 3, 4, "Hello", "ok", 45.46)
	iterator := NewSimpleIterator(l)

	for !iterator.IsDone() {
		fmt.Println(iterator.Next())
	}

}
