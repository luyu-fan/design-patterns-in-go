package _interface

// Door 门接口
type Door interface {
	Object
	BuildDoor()
	Connect(r1, r2 Room)
}
