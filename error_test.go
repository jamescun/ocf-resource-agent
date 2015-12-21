package ocf

import (
	"testing"
)

func TestErrorError(t *testing.T) {
	if s := ErrGeneric.Error(); s != "exit status 1: unspecified error" {
		t.Fatalf("expected 'exit status 1: unspecified error', got '%s'", s)
	}
}
