package types_test

import (
	"github.com/Dentrax/obscure-go/observer"
	. "github.com/Dentrax/obscure-go/types"
	"reflect"
	"testing"
)

func TestSecureInt_Inc(t *testing.T) {
	var tests = []struct {
		name  string
		in    int
		count int
		want  int
	}{
		{
			"should inc from negative",
			-77,
			77,
			0,
		},
		{
			"should inc from zero",
			0,
			77,
			77,
		},
		{
			"should inc from positive",
			77,
			77,
			154,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secInt := NewInt(tt.in).Apply()
			for i := 0; i < tt.count; i++ {
				secInt.Inc()
			}
			got := secInt.Decrypt()

			if got != tt.want {
				t.Errorf("init %v, got %+v, want %+v", secInt.GetSelf().Initialized, got, tt.want)
			}
		})
	}
}

func TestSecureInt_Dec(t *testing.T) {
	var tests = []struct {
		name  string
		in    int
		count int
		want  int
	}{
		{
			"should dec from negative",
			-77,
			77,
			-154,
		},
		{
			"should dec from zero",
			0,
			77,
			-77,
		},
		{
			"should dec from positive",
			77,
			77,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secInt := NewInt(tt.in).Apply()
			for i := 0; i < tt.count; i++ {
				secInt.Dec()
			}
			got := secInt.Decrypt()

			if got != tt.want {
				t.Errorf("init %v, got %+v, want %+v", secInt.GetSelf().Initialized, got, tt.want)
			}
		})
	}
}

func TestSecureInt_XOR(t *testing.T) {
	var tests = []struct {
		name  string
		value int
		key   int
		want  int
	}{
		{
			"should xor negative",
			-15,
			-7,
			8,
		},
		{
			"should xor zero",
			15,
			0,
			15,
		},
		{
			"should xor positive",
			15, //1111
			7,  //0111
			8,  //1000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secInt := NewInt(tt.value)
			got := secInt.GetSelf().XOR(tt.value, tt.key)

			if got != tt.want {
				t.Errorf("init %v, got %+v, want %+v", secInt.GetSelf().Initialized, got, tt.want)
			}
		})
	}
}

func TestSecureInt_Get(t *testing.T) {
	var tests = []struct {
		name  string
		value ISecureInt
		want  ISecureInt
	}{
		{
			"should get same negative",
			NewInt(-17).Apply(),
			NewInt(-17).Apply(),
		},
		{
			"should get same zero",
			NewInt(0).Apply(),
			NewInt(0).Apply(),
		},
		{
			"should get same positive",
			NewInt(17).Apply(),
			NewInt(17).Apply(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.value
			r := tt.want

			l.RandomizeKey()
			r.RandomizeKey()

			want := l.Get()
			got := r.Get()

			if got != want {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}

func TestSecureInt_Set(t *testing.T) {
	var tests = []struct {
		name  string
		value int
		want  ISecureInt
	}{
		{
			"should set zero",
			0,
			NewInt(0).Apply(),
		},
		{
			"should set negative",
			-15,
			NewInt(-15).Apply(),
		},
		{
			"should set positive",
			15,
			NewInt(15).Apply(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInt(0).Apply().Set(tt.value).Decrypt()

			if got != tt.want.Get() {
				t.Errorf("got %+v, want %+v", got, tt.want.Get())
			}
		})
	}
}

func TestSecureInt_IsEquals(t *testing.T) {
	var tests = []struct {
		name     string
		left     ISecureInt
		leftKey  int
		right    ISecureInt
		rightKey int
		want     bool
	}{
		{
			"should equals defaults",
			NewInt(0),
			111111,
			NewInt(0),
			111111,
			true,
		},
		{
			"should equals with apply",
			NewInt(77).Apply(),
			111111,
			NewInt(77).Apply(),
			111111,
			true,
		},
		{
			"should equals if no apply",
			NewInt(0).Apply(),
			111111,
			NewInt(0),
			111111,
			true,
		},
		{
			"should not equals defaults",
			NewInt(-77),
			111111,
			NewInt(77),
			111111,
			false,
		},
		{
			"should not equals with apply",
			NewInt(-77).Apply(),
			111111,
			NewInt(77).Apply(),
			111111,
			false,
		},
		{
			"should equals inc inc inc dec",
			NewInt(77).Apply().Inc().Inc().Inc().Dec(),
			111111,
			NewInt(79).Apply(),
			111111,
			true,
		},
		{
			"should equals dec dec dec inc",
			NewInt(0).Apply().Dec().Dec().Dec().Inc(),
			111111,
			NewInt(-2).Apply(),
			111111,
			true,
		},
		{
			"should not equals with different key for same value",
			NewInt(77).Apply(),
			111111,
			NewInt(77).Apply(),
			222222,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.left.SetKey(tt.leftKey)
			tt.right.SetKey(tt.rightKey)

			got := tt.left.IsEquals(tt.right)
			if got != tt.want {
				t.Errorf("got %+v, want %+v", tt.left.Decrypt(), tt.right.Decrypt())
			}
		})
	}
}

func TestSecureIntSecureInt_RandomizeKey(t *testing.T) {
	var tests = []struct {
		name  string
		value int
		after int
	}{
		{
			"should get same negative",
			-17,
			-17,
		},
		{
			"should get same zero",
			0,
			0,
		},
		{
			"should get same positive",
			17,
			17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewInt(tt.value).Apply()

			lvb := l.Get()

			l.RandomizeKey()

			lva := l.Get()

			if !reflect.DeepEqual(lva, lvb) {
				t.Errorf("must be different got %+v, want %+v", lva, lvb)
			}
		})
	}
}

func TestSecureInt_SetKey(t *testing.T) {
	tests := []struct {
		name string
		key  int
	}{
		{
			"should get same positive",
			123456789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewInt(0).Apply()
			l.SetKey(tt.key)

			if l.GetSelf().Key != tt.key {
				t.Errorf("got %+v, want %+v", l.GetSelf().Key, tt.key)
			}
		})
	}
}

func TestSecureInt_Decrypt(t *testing.T) {
	tests := []struct {
		name string
		data *SecureInt
		want int
	}{
		{
			"should not initialized",
			&SecureInt{
				Key:         0,
				RealValue:   0,
				Initialized: false,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.data.Decrypt()

			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSecureInt_AddWatcher(t *testing.T) {
	tests := []struct {
		name      string
		fakeValue int
		watcher   observer.Observer
	}{
		{
			"should attach",
			7,
			observer.CreateWatcher("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := NewInt(0).Apply()

			data.AddWatcher(tt.watcher)

			if len(data.GetSelf().Observers) == 0 {
				t.Errorf("len(tt.data.Observers) == 0: got %+v, want %+v", 0, len(data.GetSelf().Observers))
			}

			if data.GetSelf().HackDetecting != true {
				t.Errorf("tt.data.HackDetecting != true: got %+v, want %+v", data.GetSelf().HackDetecting, true)
			}

			data.Set(tt.fakeValue)

			if data.GetSelf().FakeValue != tt.fakeValue {
				t.Errorf("tt.data.FakeValue != 7: got %+v, want %+v", data.GetSelf().FakeValue, tt.fakeValue)
			}

			data.Inc()

			if data.GetSelf().FakeValue != tt.fakeValue+1 {
				t.Errorf("tt.data.FakeValue != fakeValue + 1: got %+v, want %+v", data.GetSelf().FakeValue, tt.fakeValue+1)
			}

			data.Dec()
			data.Dec()

			if data.GetSelf().FakeValue != tt.fakeValue-1 {
				t.Errorf("tt.data.FakeValue != fakeValue - 1: got %+v, want %+v", data.GetSelf().FakeValue, tt.fakeValue-1)
			}

			data.GetSelf().FakeValue = 999

			data.Decrypt()

			if len(data.GetSelf().Notifies) == 0 {
				t.Errorf("len(data.GetSelf().Notifies) == 0: got %+v, want %+v", len(data.GetSelf().Notifies), 1)
			}
		})
	}
}

func Benchmark_PrimitiveInt(b *testing.B) {
	o := 37

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		i := 17
		i++
		i--

		if i != o {
			o--
		}
	}
}

func Benchmark_SecureInt(b *testing.B) {
	o := NewInt(37).Apply()
	i := NewInt(17).Apply()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		i.Inc()
		i.Dec()

		if !i.IsEquals(o) {
			o.Dec()
		}
	}
}
