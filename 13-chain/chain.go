package _3_chain

import "fmt"

type requestType uint8

const (
	// BOTTOM_REQ 低层请求
	BOTTOM_REQ requestType = iota + 1
	// MIDDLE_RRQ 中层请求
	MIDDLE_RRQ
	// TOP_REQ 顶层请求
	TOP_REQ
)

// RequestHandler 请求处理对象接口
type RequestHandler interface {
	// SetSuccessor 设置后继请求对象
	SetSuccessor(handler RequestHandler)
	// HandleRequest 处理请求
	HandleRequest(request CommonRequest)
}

// CommonRequest 通用请求接口
type CommonRequest interface {
	GetRequestType() requestType
}

type topHandler struct {
	successor RequestHandler
}

func (th *topHandler) SetSuccessor(handler RequestHandler) {
	th.successor = handler
}

func (th *topHandler) HandleRequest(request CommonRequest) {
	if request.GetRequestType() == TOP_REQ {
		fmt.Println("I am topHandler, I caught the request")
	} else {
		fmt.Println("I am topHandler, Request type is out of range. Error")
	}
}

type middleHandler struct {
	successor RequestHandler
}

func (mh *middleHandler) SetSuccessor(handler RequestHandler) {
	mh.successor = handler
}

func (mh *middleHandler) HandleRequest(request CommonRequest) {
	if request.GetRequestType() == MIDDLE_RRQ {
		fmt.Println("I am middleHandler, I caught the request")
	} else {
		fmt.Println("I am middleHandler, pop the request")
		mh.successor.HandleRequest(request)
	}
}

type bottomHandler struct {
	successor RequestHandler
}

func (bh *bottomHandler) SetSuccessor(handler RequestHandler) {
	bh.successor = handler
}

func (bh *bottomHandler) HandleRequest(request CommonRequest) {
	if request.GetRequestType() == BOTTOM_REQ {
		fmt.Println("I am bottomHandler, I caught the request")
	} else {
		fmt.Println("I am bottomHandler, pop the request")
		bh.successor.HandleRequest(request)
	}
}

type basicRequest struct {
	reqType requestType
}

func (br *basicRequest) GetRequestType() requestType {
	return br.reqType
}

func NewRequest(reqType requestType) CommonRequest {
	return &basicRequest{reqType: reqType}
}

// 具体的请求处理对象 每个对象都有自己独有的方法 这里省略

func NewBottomHandler() RequestHandler {
	return &bottomHandler{}
}

func NewMiddleHandler() RequestHandler {
	return &middleHandler{}
}

func NewTopHandler() RequestHandler {
	return &topHandler{}
}
