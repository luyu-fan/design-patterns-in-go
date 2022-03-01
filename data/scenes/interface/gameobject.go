package _interface

type GameObject interface {
	Object
	Clone() GameObject
	Initialize(v ...interface{})
}
