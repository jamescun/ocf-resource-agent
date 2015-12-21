package ocf

import (
	"testing"
)

var testEnv = Env{
	"invalid",
	"OCF_RESOURCE_INSTANCE=foobar",
	"OCF_RESKEY_CRM_meta_notify_start_uname=example.org",
	"OCF_RESKEY_CRM_meta_notify_master_uname=m1.example.org m2.example.org",
}

func TestEnvGet(t *testing.T) {
	if s := testEnv.Get("OCF_RESOURCE_INSTANCE"); s != "foobar" {
		t.Fatalf("expected 'foobar', got '%s'", s)
	} else if s := testEnv.Get("none"); s != "" {
		t.Fatalf("expected '', got '%s'", s)
	}
}

func BenchmarkEnvGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testEnv.Get("OCF_RESOURCE_INSTANCE")
	}
}

func TestCountFields(t *testing.T) {
	if c := testEnv.CountFields("OCF_RESKEY_CRM_"); c != 2 {
		t.Fatal("expected 2 fields, got", c)
	}
}

func BenchmarkCountFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testEnv.CountFields("OCF_RESKEY_CRM_")
	}
}

func TestEnvFields(t *testing.T) {
	f := testEnv.Fields("OCF_RESKEY_CRM_")
	if len(f) != 2 {
		t.Fatal("expected 2 fields, got", len(f))
	}

	if v, ok := f["meta_notify_start_uname"]; ok {
		if v != "example.org" {
			t.Fatalf("expected meta_notify_start_uname 'example.org', got '%s'", v)
		}
	} else {
		t.Fatal("expected key meta_notify_start_uname")
	}
}

func BenchmarkEnvFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testEnv.Fields("OCF_RESKEY_CRM_")
	}
}
