package crm

import (
	"testing"
)

var testCRM = CRM{
	Property:   "ocf",
	MasterPath: "test/crm_master",
	AttrPath:   "test/crm_attribute",
}

func TestNew(t *testing.T) {
	crm := New("replication")
	if crm.Property != "replication" {
		t.Fatal("expected property 'replication', got '%s'", crm.Property)
	}
}

func TestCRMGet(t *testing.T) {
	v, err := testCRM.Get("master")
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if v != "example.org" {
		t.Fatalf("expected 'example.org', got '%s'", v)
	}
}

func TestCRMSet(t *testing.T) {
	err := testCRM.Set("master", "me")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestCRMSetScore(t *testing.T) {
	err := testCRM.SetScore(6379)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestCRMGetScore(t *testing.T) {
	n, err := testCRM.GetScore()
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if n != 6379 {
		t.Fatal("expected score 6379, got", n)
	}
}

func TestCRMDeleteScore(t *testing.T) {
	err := testCRM.DeleteScore()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}
