package types

import (
	"fmt"
	obs "github.com/Dentrax/obscure-go/observer"
	"math/rand"
	"time"
)

const KEY int = 54343

type ISecureInt interface {
	Apply() ISecureInt
	AddWatcher(obs obs.Observer)
	SetKey(int)
	Inc() ISecureInt
	Dec() ISecureInt
	Set(int) ISecureInt
	Get() int
	GetSelf() *SecureInt
	Decrypt() int
	RandomizeKey()
	IsEquals(ISecureInt) bool
}

type SecureInt struct {
	obs.Observable
	Key           int
	RealValue     int
	FakeValue     int
	Initialized   bool
	HackDetecting bool
}

func NewInt(value int) ISecureInt {
	s := &SecureInt{
		Key:           KEY,
		RealValue:     value,
		FakeValue:     value,
		Initialized:   false,
		HackDetecting: false,
	}

	s.Apply()

	return s
}

func (i *SecureInt) Apply() ISecureInt {
	if !i.Initialized {
		i.RealValue = i.XOR(i.RealValue, i.Key)
		i.Initialized = true
	}

	return i
}

func (i *SecureInt) AddWatcher(obs obs.Observer) {
	i.AddObserver(obs)
	i.HackDetecting = true
}

func (i *SecureInt) SetKey(key int) {
	i.Key = key
}

func (i *SecureInt) RandomizeKey() {
	rand.Seed(time.Now().UnixNano())

	i.RealValue = i.Decrypt()
	i.Key = rand.Intn(int(^uint(0) >> 1))
	i.RealValue = i.XOR(i.RealValue, i.Key)
}

func (i *SecureInt) XOR(value int, Key int) int {
	return value ^ Key
}

func (i *SecureInt) Get() int {
	return i.Decrypt()
}

func (i *SecureInt) GetSelf() *SecureInt {
	return i
}

func (i *SecureInt) Set(value int) ISecureInt {
	i.RealValue = i.XOR(value, i.Key)

	if i.HackDetecting {
		i.FakeValue = value
	}

	return i
}

func (i *SecureInt) Decrypt() int {
	if !i.Initialized {
		i.Initialized = false
		i.FakeValue = 0
		i.RealValue = i.XOR(0, 0)
		i.Key = KEY

		return 0
	}

	res := i.XOR(i.RealValue, i.Key)

	if i.HackDetecting && res != i.FakeValue {
		i.NotifyAll(fmt.Sprintf("hack attempt: %d", i.FakeValue))
	}

	return res
}

func (i *SecureInt) Inc() ISecureInt {
	next := i.Decrypt() + 1

	i.RealValue = i.XOR(next, i.Key)

	if i.HackDetecting {
		i.FakeValue = next
	}

	return i
}

func (i *SecureInt) Dec() ISecureInt {
	next := i.Decrypt() - 1

	i.RealValue = i.XOR(i.Decrypt()-1, i.Key)

	if i.HackDetecting {
		i.FakeValue = next
	}

	return i
}

func (i *SecureInt) IsEquals(o ISecureInt) bool {
	if i.Key != o.GetSelf().Key {
		return i.XOR(i.RealValue, i.Key) == i.XOR(o.GetSelf().RealValue, o.GetSelf().Key)
	}

	return i.RealValue == o.GetSelf().RealValue
}
