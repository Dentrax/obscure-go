package types

import (
	"math/rand"
	"time"
)

const KEY int = 54343

type ISecureInt interface {
	Apply() ISecureInt
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
	key         int
	RealValue   int
	fakeValue   int
	Initialized bool
}

func NewInt(value int) ISecureInt {
	s := &SecureInt{
		key:         KEY,
		RealValue:   value,
		fakeValue:   value,
		Initialized: false,
	}

	s.Apply()

	return s
}

func (i *SecureInt) Apply() ISecureInt {
	if !i.Initialized {
		i.RealValue = i.XOR(i.RealValue, i.key)
		i.Initialized = true
	}

	return i
}

func (i *SecureInt) SetKey(key int) {
	i.key = key
}

func (i *SecureInt) RandomizeKey() {
	rand.Seed(time.Now().UnixNano())

	i.RealValue = i.Decrypt()
	i.key = rand.Intn(int(^uint(0) >> 1))
	i.RealValue = i.XOR(i.RealValue, i.key)
}

func (i *SecureInt) XOR(value int, key int) int {
	return value ^ key
}

//Why not
func (i *SecureInt) Get() int {
	return i.Decrypt()
}

func (i *SecureInt) GetSelf() *SecureInt {
	return i
}

func (i *SecureInt) Set(value int) ISecureInt {
	i.RealValue = i.XOR(value, i.key)
	return i
}

func (i *SecureInt) Decrypt() int {
	if !i.Initialized {
		i.Initialized = false
		i.fakeValue = 0
		i.RealValue = i.XOR(0, 0)
		i.key = KEY

		return 0
	}

	return i.XOR(i.RealValue, i.key)
}

func (i *SecureInt) Inc() ISecureInt {
	i.RealValue = i.XOR(i.Decrypt()+1, i.key)
	return i
}

func (i *SecureInt) Dec() ISecureInt {
	i.RealValue = i.XOR(i.Decrypt()-1, i.key)
	return i
}

func (i *SecureInt) IsEquals(o ISecureInt) bool {
	if i.key != o.GetSelf().key {
		return i.XOR(i.RealValue, i.key) == i.XOR(o.GetSelf().RealValue, o.GetSelf().key)
	}

	return i.RealValue == o.GetSelf().RealValue
}
