package ocf

import (
	"testing"
)

var testContext = &Context{
	Keys: map[string]string{
		"meta_notify_type":         "pre",
		"meta_notify_operation":    "promote",
		"meta_notify_start_uname":  "example",
		"meta_notify_master_uname": "foo bar",
		"agent_internal_timeout":   "3600",
	},
}

func TestContextReadNotifyType(t *testing.T) {
	testContext.readNotifyType()
	if testContext.Notify != "pre-promote" {
		t.Fatalf("expected Notify 'pre-promote', got '%s'", testContext.Notify)
	}
}

func BenchmarkReadNotifyType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testContext.readNotifyType()
	}
}

func TestContextString(t *testing.T) {
	if s := testContext.String("meta_notify_start_uname"); s != "example" {
		t.Fatalf("expected 'example', got '%s'", s)
	} else if s := testContext.String("foo"); s != "" {
		t.Fatalf("expected empty string, got '%s'", s)
	}
}

func BenchmarkContextString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testContext.String("meta_notify_start_uname")
	}
}

func TestContextStringArray(t *testing.T) {
	a := testContext.StringArray("meta_notify_master_uname")
	if len(a) != 2 {
		t.Fatal("expected 2 strings, got", len(a))
	} else if a[0] != "foo" {
		t.Fatalf("expected a[0] 'foo', got '%s'", a[0])
	} else if a[1] != "bar" {
		t.Fatalf("expected a[1] 'bar', got '%s'", a[1])
	}
}

func BenchmarkContextStringArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testContext.StringArray("meta_notify_master_uname")
	}
}

func TestContextInt(t *testing.T) {
	if v := testContext.Int("agent_internal_timeout"); v != 3600 {
		t.Fatal("expected 3600, got", v)
	} else if v := testContext.Int("foo"); v != 0 {
		t.Fatal("expected 0, got", v)
	}
}

func BenchmarkContextInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testContext.Int("agent_internal_timeout")
	}
}

func TestContextInt64(t *testing.T) {
	if v := testContext.Int64("agent_internal_timeout"); v != 3600 {
		t.Fatal("expected 3600, got", v)
	} else if v := testContext.Int64("foo"); v != 0 {
		t.Fatal("expected 0, got", v)
	}
}

func BenchmarkContextInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testContext.Int64("agent_internal_timeout")
	}
}
