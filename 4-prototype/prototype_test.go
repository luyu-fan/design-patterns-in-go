package __prototype

import "testing"

func TestClientGetPrototype(t *testing.T) {
	client := Client{}
	client.GetPrototype()
}

func TestClientGetPrototypeAndCopy(t *testing.T) {
	client := Client{}
	client.GetPrototypeAndCopy()
}

func BenchmarkClientGetPrototype(b *testing.B) {
	client := Client{}
	for i := 0; i < b.N; i++ {
		client.GetPrototype()
	}
}

func BenchmarkClientGetPrototypeAndCopy(b *testing.B) {
	client := Client{}
	for i := 0; i < b.N; i++ {
		client.GetPrototypeAndCopy()
	}
}
