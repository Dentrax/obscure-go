package main

import (
	"math/rand"
	"time"
)

const KEY int = 54343

type SecureInt struct {
	key         int
	realValue   int
	fakeValue   int
	initialized bool
}

func NewSecureInt(value int) *SecureInt {
	return &SecureInt{
		key:         KEY,
		realValue:   value,
		fakeValue:   value,
		initialized: false,
	}
}

func (i *SecureInt) Apply() *SecureInt {
	if !i.initialized {
		i.realValue = i.XOR(i.realValue, i.key)
		i.initialized = true
	}
	return i
}

func (i *SecureInt) SetKey(key int) {
	i.key = key
}

func (i *SecureInt) RandomizeKey() {
	rand.Seed(time.Now().UnixNano())

	i.realValue = i.Decrypt()
	i.key = rand.Intn(int(^uint(0) >> 1))
	i.realValue = i.XOR(i.realValue, i.key)
}

func (i *SecureInt) XOR(value int, key int) int {
	return value ^ key
}

//Why not
func (i *SecureInt) Get() int {
	return i.Decrypt()
}

func (i *SecureInt) Set(value int) *SecureInt {
	i.realValue = i.XOR(value, i.key)
	return i
}

func (i *SecureInt) Decrypt() int {
	if !i.initialized {
		i.initialized = false
		i.fakeValue = 0
		i.realValue = i.XOR(0, 0)
		i.key = KEY
		return 0
	}

	return i.XOR(i.realValue, i.key)
}

func (i *SecureInt) Inc() *SecureInt {
	i.realValue = i.XOR(i.Decrypt()+1, i.key)
	return i
}

func (i *SecureInt) Dec() *SecureInt {
	i.realValue = i.XOR(i.Decrypt()-1, i.key)
	return i
}

func (i *SecureInt) IsEquals(o *SecureInt) bool {
	if i.key != o.key {
		return i.XOR(i.realValue, i.key) == i.XOR(o.realValue, o.key)
	}

	return i.realValue == o.realValue
}
