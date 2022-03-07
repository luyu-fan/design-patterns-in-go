package _0_state

import (
	"fmt"
	"math/rand"
)

type LifeState interface {
	Live() LifeState
	GetState() string
}

type Person interface {
	Living()
}

type richMan struct {
	state LifeState
}

func (pm *richMan) Living() {
	for pm.state.GetState() != "dead" {
		pm.state = pm.state.Live()
	}
}

type youngState struct {
	age uint
}

func (bs *youngState) Study() {
	fmt.Println("Everything is curious")
}

func (bs *youngState) Work() {
	fmt.Println("No, babe does not work, but a little for the poor")
}

func (bs *youngState) Play() {
	fmt.Println("Haha, mostly, babe plays carefree")
}

func (bs *youngState) Live() LifeState {
	for bs.age < 10 {
		bs.Work()
		bs.Study()
		bs.Play()
		bs.age += 1
	}
	return &manState{}
}

func (bs *youngState) GetState() string {
	return "young"
}

type manState struct {
	money int
	age   int
}

func (ms *manState) GetState() string {
	return "man"
}

func (ms *manState) Live() LifeState {
	ms.age = 20
	for ms.money <= 1000*100000 && ms.age < 65 {
		if rand.Intn(10) < 6 {
			ms.Study()
		}
		ms.Work()
	}
	fmt.Println("Money: ", ms.money, "Age: ", ms.age)
	return &oldState{}
}

func (ms *manState) Work() {
	fmt.Println("Man has to earn money")
	ms.money += rand.Intn(50 * 10000)
	ms.age++

}

func (ms *manState) Study() {
	fmt.Println("Man has to learn")
}

type oldState struct {
	age int
}

func (os *oldState) GetState() string {
	return "old"
}

func (os *oldState) Enjoy() {
	os.age++
	if rand.Intn(10) < 6 {
		fmt.Println("Enjoying time")
	} else {
		fmt.Println("Stile working")
	}

}

func (os *oldState) Live() LifeState {
	os.age = 65
	totalAge := rand.Intn(50) + 65
	for os.age < totalAge {
		os.Enjoy()
	}
	fmt.Println("Dead at ", os.age)
	return &deadState{}
}

type deadState struct {
}

func (ds *deadState) GetState() string {
	return "dead"
}

func (ds *deadState) Live() LifeState {
	return nil
}

func NewMan() Person {
	return &richMan{
		state: &youngState{age: 1},
	}
}
