package __decorator

import "testing"

func TestDecorator(t *testing.T) {

	tv := textView{
		text: "这是一段文本",
	}

	bd := borderDecorator{
		&visualDecorator{&tv},
	}

	bd.Draw()
	bd.Resize()

}
