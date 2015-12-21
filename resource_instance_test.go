package ocf

import (
	"strconv"
	"testing"
)

func TestParseResourceInstance(t *testing.T) {
	ri, err := ParseResourceInstance(" foobar:1 ")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if ri.Name != "foobar" {
		t.Fatalf("expected name 'foobar', got '%s'", ri.Name)
	} else if ri.Clone != 1 {
		t.Fatal("expected clone 1, got", ri.Clone)
	}
}

func TestParseResourceInstanceNoClone(t *testing.T) {
	ri, err := ParseResourceInstance(" foobar ")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if ri.Name != "foobar" {
		t.Fatalf("expected name 'foobar', got '%s'", ri.Name)
	} else if ri.Clone != 0 {
		t.Fatal("expected clone 0, got", ri.Clone)
	}
}

func TestParseResourceInstanceEmpty(t *testing.T) {
	_, err := ParseResourceInstance("")
	if err != ErrUnimplemented {
		t.Fatalf("expected ocf.Error (ErrUnimplemented), got %T", err)
	}
}

func TestParseResourceInstanceInvalidClone(t *testing.T) {
	_, err := ParseResourceInstance("foobar:NOPE")
	if _, ok := err.(*strconv.NumError); !ok {
		t.Fatalf("expected *strconv.NumError, got %T", err)
	}
}

func BenchmarkParseResourceInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseResourceInstance("foobar:1")
	}
}
