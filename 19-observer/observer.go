package _9_observer

import (
	"fmt"
	"time"
)

type signature uint8

// Subject 目标接口
type Subject interface {
	GetTopic() string
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

// Observer 观察者接口
type Observer interface {
	// GetSignature 观察者身份
	GetSignature() signature
	// Update 扩展的更新接口使得观察者能够同时观测多个对象
	Update(subject Subject)
}

// baseSubject 目标的抽象类
type baseSubject struct {
	topic string
	// 用于维护映射关系的简单map
	obs map[signature]Observer
}

func (bs *baseSubject) GetTopic() string {
	return bs.topic
}

func (bs *baseSubject) Attach(ob Observer) {
	bs.obs[ob.GetSignature()] = ob
}

func (bs *baseSubject) Detach(ob Observer) {
	delete(bs.obs, ob.GetSignature())
}

func (bs *baseSubject) Notify() {
	// ... default implementation
	for i := range bs.obs {
		bs.obs[i].Update(bs)
	}
}

type ClockTimer struct {
	baseSubject
}

func (ct *ClockTimer) Notify() {
	// Override
	for i := range ct.obs {
		ct.obs[i].Update(ct)
	}
}

func (ct *ClockTimer) Tick() {
	ct.Notify()
}

func (ct *ClockTimer) GetHour() int {
	now := time.Now()
	return now.Hour()
}

func (ct *ClockTimer) GetMinute() int {
	now := time.Now()
	return now.Minute()
}

func (ct *ClockTimer) GetSecond() int {
	now := time.Now()
	return now.Second()
}

type DigitalClock struct {
	sig     signature
	subject Subject
}

func (dc *DigitalClock) GetSignature() signature {
	return dc.sig
}

func (dc *DigitalClock) Update(sub Subject) {
	switch sub.(type) {
	case *ClockTimer:
		{
			// 获取目标的信息
			sub := sub.(*ClockTimer)
			hour := sub.GetHour()
			min := sub.GetMinute()
			sec := sub.GetSecond()
			fmt.Println("Digital Clock Observing:", hour, min, sec)
		}

	default:
		fmt.Println("Wrong Type")
	}
}
