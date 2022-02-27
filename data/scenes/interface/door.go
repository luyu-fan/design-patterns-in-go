package _interface

// Door 门接口
type Door interface {
	GameObject
	BuildDoor()
	Connect(r1, r2 Room)
}
