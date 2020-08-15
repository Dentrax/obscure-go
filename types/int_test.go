package types_test

import (
	. "github.com/Dentrax/obscure-go/types"
	"testing"
)

func Test_Inc(t *testing.T) {
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

func Test_Dec(t *testing.T) {
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

func Test_XOR(t *testing.T) {
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

func Test_Get(t *testing.T) {
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

func Test_Set(t *testing.T) {
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
			want := NewInt(0).Apply().Set(tt.value).Decrypt()
			got := tt.value

			if got != want {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}

func Test_IsEquals(t *testing.T) {
	var tests = []struct {
		name  string
		left  ISecureInt
		right ISecureInt
		want  bool
	}{
		{
			"should equals defaults",
			NewInt(0),
			NewInt(0),
			true,
		},
		{
			"should equals with apply",
			NewInt(77).Apply(),
			NewInt(77).Apply(),
			true,
		},
		{
			"should equals if no apply",
			NewInt(0).Apply(),
			NewInt(0),
			true,
		},
		{
			"should not equals defaults",
			NewInt(-77),
			NewInt(77),
			false,
		},
		{
			"should not equals with apply",
			NewInt(-77).Apply(),
			NewInt(77).Apply(),
			false,
		},
		{
			"should equals inc inc inc dec",
			NewInt(77).Apply().Inc().Inc().Inc().Dec(),
			NewInt(79).Apply(),
			true,
		},
		{
			"should equals dec dec dec inc",
			NewInt(0).Apply().Dec().Dec().Dec().Inc(),
			NewInt(-2).Apply(),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.left.IsEquals(tt.right)
			if got != tt.want {
				t.Errorf("got %+v, want %+v", tt.left.Decrypt(), tt.right.Decrypt())
			}
		})
	}
}

func Test_RandomizeKey(t *testing.T) {
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

			lvb := l.GetSelf().RealValue

			l.RandomizeKey()

			lva := l.GetSelf().RealValue

			if lva == lvb {
				t.Errorf("must be different got %+v, want %+v", lva, lvb)
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
