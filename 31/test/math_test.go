package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("do some setup")
	m.Run()
	fmt.Println("do some cleanup")
}

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func TestAbs_2(t *testing.T) {
	tests := []struct {
		x    float64
		want float64
	}{
		{-0.3, 0.3},
		{-2, 2},
		{-3.1, 3.1},
		{5, 5},
	}

	for _, tt := range tests {
		if got := Abs(tt.x); got != tt.want {
			t.Errorf("Abs() = %f, want %v", got, tt.want)
		}
	}
}

func TestAbs_3(t *testing.T) {
	tests := []struct {
		x    float64
		want float64
	}{
		{-0.3, 0.3},
		{-2, 2},
		{-3.1, 3.1},
		{5, 5},
	}

	for _, tt := range tests {
		got := Abs(tt.x)
		assert.Equal(t, got, tt.want)
	}
}

func TestMax(t *testing.T) {
	got := Max(1, 2)
	if got != 2 {
		t.Errorf("Max(1, 2) = %f; want 2", got)
	}
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt()
	}
}

func TestMin(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandInt(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandInt(); got != tt.want {
				t.Errorf("RandInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
