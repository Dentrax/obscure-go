package types

import (
	"github.com/Dentrax/obscure-go/observer"
	"reflect"
	"testing"
)

func TestSecureString_XOR(t *testing.T) {
	tests := []struct {
		name  string
		value string
		key   int
		want  string
	}{
		{
			"should xor empty",
			"",
			1,
			"",
		},
		{
			"should xor string with zero",
			"Hello World",
			0,
			"Hello World",
		},
		{
			"should xor string",
			"Hello World",
			7,
			"Obkkh'Phukc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secStr := NewString(tt.value)
			got := secStr.GetSelf().XOR([]rune(tt.value), tt.key)

			if len(got) != len(tt.want) {
				t.Errorf("len(got) != len(tt.want) %v, got %+v, want %+v", secStr.GetSelf().Initialized, string(got), string(tt.want))
			}

			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("v != got[i] %v, got %+v, want %+v", secStr.GetSelf().Initialized, string(got), string(tt.want))
			}
		})
	}
}

func TestSecureString_SetKey(t *testing.T) {
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
			l := NewString("").Apply()
			l.SetKey(tt.key)

			if l.GetSelf().Key != tt.key {
				t.Errorf("got %+v, want %+v", l.GetSelf().Key, tt.key)
			}
		})
	}
}

func TestSecureString_RandomizeKey(t *testing.T) {
	var tests = []struct {
		name  string
		value string
		after string
	}{
		{
			"should get same foo",
			"foo",
			"foo",
		},
		{
			"should get same empty",
			"",
			"",
		},
		{
			"should get same bar",
			"bar",
			"bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewString(tt.value).Apply()

			lvb := l.Get()

			l.RandomizeKey()

			lva := l.Get()

			if len(lva) != len(lvb) {
				t.Errorf("must be different len() got %+v, want %+v", string(lva), string(lvb))
			}

			if !reflect.DeepEqual(lva, string(lvb)) {
				t.Errorf("must be different got %+v, want %+v", string(lva), string(lvb))
			}
		})
	}
}

func TestSecureString_IsEquals(t *testing.T) {
	var tests = []struct {
		name     string
		left     ISecureString
		leftKey  int
		right    ISecureString
		rightKey int
		want     bool
	}{
		{
			"should equals defaults",
			NewString(""),
			111111,
			NewString(""),
			111111,
			true,
		},
		{
			"should equals with apply",
			NewString("foo").Apply(),
			111111,
			NewString("foo").Apply(),
			111111,
			true,
		},
		{
			"should equals if no apply",
			NewString("").Apply(),
			111111,
			NewString(""),
			111111,
			true,
		},
		{
			"should not equals defaults",
			NewString("foo"),
			111111,
			NewString("bar"),
			111111,
			false,
		},
		{
			"should not equals with apply",
			NewString("foo").Apply(),
			111111,
			NewString("bar").Apply(),
			111111,
			false,
		},
		{
			"should not equals with different key for same value",
			NewString("foo").Apply(),
			111111,
			NewString("foo").Apply(),
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

func TestSecureString_Decrypt(t *testing.T) {
	tests := []struct {
		name string
		data *SecureString
		want []rune
	}{
		{
			"should not initialized",
			&SecureString{
				Key:         0,
				RealValue:   []rune(""),
				Initialized: false,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.data.Decrypt()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSecureString_Set(t *testing.T) {
	var tests = []struct {
		name    string
		value   string
		want    ISecureString
		wantErr bool
	}{
		{
			"should set empty",
			"",
			NewString("").Apply(),
			false,
		},
		{
			"should set foo",
			"foo",
			NewString("foo").Apply(),
			false,
		},
		{
			"should set bar",
			"bar",
			NewString("foo").Apply(),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewString("").Apply().Set(tt.value).Get()

			if (got != tt.want.Get()) != tt.wantErr {
				t.Errorf("got %+v, want %+v", got, tt.want.Get())
			}
		})
	}
}

func TestSecureString_AddWatcher(t *testing.T) {
	tests := []struct {
		name      string
		fakeValue string
		watcher   observer.Observer
	}{
		{
			"should attach",
			"foo",
			observer.CreateWatcher("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := NewString("").Apply()

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

			data.GetSelf().FakeValue = "bar"

			data.Decrypt()

			if len(data.GetSelf().Notifies) == 0 {
				t.Errorf("len(data.GetSelf().Notifies) == 0: got %+v, want %+v", len(data.GetSelf().Notifies), 1)
			}
		})
	}
}

func Benchmark_PrimitiveString(b *testing.B) {
	o := "foo"

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		i := "bar"
		i = "baz"
		i = "qux"

		if i != o {
			o = "quux"
		}
	}
}

func Benchmark_SecureString(b *testing.B) {
	o := NewString("foo").Apply()
	i := NewString("bar").Apply()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		i.Set("baz")
		i.Set("qux")

		if !i.IsEquals(o) {
			o.Set("quux")
		}
	}
}
