package _3_chain

import "testing"

func TestChain(t *testing.T) {

	bh := NewBottomHandler()
	mh := NewMiddleHandler()
	th := NewTopHandler()

	bh.SetSuccessor(mh)
	mh.SetSuccessor(th)

	req1 := NewRequest(1)
	req2 := NewRequest(2)
	req3 := NewRequest(3)
	req4 := NewRequest(4)

	bh.HandleRequest(req1)
	bh.HandleRequest(req2)
	bh.HandleRequest(req3)
	bh.HandleRequest(req4)

	mh.HandleRequest(req1)
	mh.HandleRequest(req2)
	mh.HandleRequest(req3)
	mh.HandleRequest(req4)

	th.HandleRequest(req1)
	th.HandleRequest(req2)
	th.HandleRequest(req3)
	th.HandleRequest(req4)

}
