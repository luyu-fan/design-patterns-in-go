package _7_mediator

type DialogMediator interface {
	ShowDialog()
	WidgetChanged(w Widget)
}

type Widget interface {
}
