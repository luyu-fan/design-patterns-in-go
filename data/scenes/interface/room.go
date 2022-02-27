package _interface

// Room 房间接口
type Room interface {
	GameObject
	BuildRoom()
	AddDoor(d Door)
	AddWall(w Wall)
}
