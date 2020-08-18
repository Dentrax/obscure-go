package types

import (
	obs "github.com/Dentrax/obscure-go/observer"
	"math/rand"
	"time"
)

type ISecureString interface {
	Apply() ISecureString
	AddWatcher(obs obs.Observer)
	SetKey(int)
	Set(string) ISecureString
	Get() string
	GetSelf() *SecureString
	Decrypt() []rune
	RandomizeKey()
	IsEquals(ISecureString) bool
}

type SecureString struct {
	obs.Observable
	Key           int
	RealValue     []rune
	FakeValue     string
	Initialized   bool
	HackDetecting bool
}

func NewString(value string) ISecureString {
	s := &SecureString{
		Key:           KEY,
		RealValue:     []rune(value),
		FakeValue:     value,
		Initialized:   false,
		HackDetecting: false,
	}

	s.Apply()

	return s
}

func (i *SecureString) Apply() ISecureString {
	if !i.Initialized {
		i.RealValue = i.XOR(i.RealValue, i.Key)
		i.Initialized = true
	}

	return i
}

func (i *SecureString) AddWatcher(obs obs.Observer) {
	i.AddObserver(obs)
	i.HackDetecting = true
}

func (i *SecureString) SetKey(key int) {
	i.Key = key
}

func (i *SecureString) RandomizeKey() {
	rand.Seed(time.Now().UnixNano())

	i.RealValue = i.Decrypt()
	i.Key = rand.Intn(int(^uint(0) >> 1))
	i.RealValue = i.XOR(i.RealValue, i.Key)
}

func (i *SecureString) XOR(value []rune, key int) []rune {
	res := make([]rune, len(value))

	for i, v := range value {
		res[i] = v ^ int32(key)
	}

	return res
}

func (i *SecureString) Get() string {
	return string(i.Decrypt())
}

func (i *SecureString) GetSelf() *SecureString {
	return i
}

func (i *SecureString) Set(value string) ISecureString {
	i.RealValue = i.XOR([]rune(value), i.Key)

	if i.HackDetecting {
		i.FakeValue = value
	}

	return i
}

func (i *SecureString) Decrypt() []rune {
	if !i.Initialized {
		i.Key = KEY
		i.FakeValue = ""
		i.RealValue = i.XOR(nil, 0)
		i.Initialized = false

		return nil
	}

	res := i.XOR(i.RealValue, i.Key)

	if i.HackDetecting && string(res) != i.FakeValue {
		i.NotifyAll("hack")
	}

	return res
}

func (i *SecureString) IsEquals(o ISecureString) bool {
	if i.Key != o.GetSelf().Key {
		return string(i.XOR(i.RealValue, i.Key)) == string(i.XOR(o.GetSelf().RealValue, o.GetSelf().Key))
	}

	return string(i.RealValue) == string(o.GetSelf().RealValue)
}
