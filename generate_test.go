package joeusername

import "testing"

func TestGenerate(t *testing.T) {
	var want string
	if got := Generate(1); len(got) <= 0 {
		t.Errorf("Generate() = %T, want %T", got, want)
	}
}
