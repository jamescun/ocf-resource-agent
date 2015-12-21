package ocf

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

type resource struct {
	lastRun string
}

func (r *resource) Start(c *Context) error   { r.lastRun = "start"; return nil }
func (r *resource) Stop(c *Context) error    { r.lastRun = "stop"; return nil }
func (r *resource) Monitor(c *Context) error { r.lastRun = "monitor"; return nil }
func (r *resource) Notify(c *Context) error  { r.lastRun = "notify"; return nil }
func (r *resource) Promote(c *Context) error { r.lastRun = "promote"; return nil }
func (r *resource) Demote(c *Context) error  { r.lastRun = "demote"; return nil }
func (r *resource) Metadata() error          { r.lastRun = "metadata"; return nil }

var testResource = &resource{}
var testAgent = New(testResource)

func TestNew(t *testing.T) {
	a := New(testResource)
	if a == nil {
		t.Fatal("expected Agent, got nil")
	} else if a.r == nil {
		t.Fatal("Agent does not have Resource")
	}
}

func TestAgentRunStart(t *testing.T) {
	err := testAgent.run("start", testContext)
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if testResource.lastRun != "start" {
		t.Fatal("action start was not invoked")
	}
}

func TestAgentRunStop(t *testing.T) {
	err := testAgent.run("stop", testContext)
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if testResource.lastRun != "stop" {
		t.Fatal("action stop was not invoked")
	}
}

func TestAgentRunMonitor(t *testing.T) {
	err := testAgent.run("monitor", testContext)
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if testResource.lastRun != "monitor" {
		t.Fatal("action monitor was not invoked")
	}
}
func TestAgentRunNotify(t *testing.T) {
	err := testAgent.run("notify", testContext)
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if testResource.lastRun != "notify" {
		t.Fatal("action notify was not invoked")
	}
}
func TestAgentRunPromote(t *testing.T) {
	err := testAgent.run("promote", testContext)
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if testResource.lastRun != "promote" {
		t.Fatal("action promote was not invoked")
	}
}

func TestAgentRunDemote(t *testing.T) {
	err := testAgent.run("demote", testContext)
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if testResource.lastRun != "demote" {
		t.Fatal("action demote was not invoked")
	}
}

func TestAgentRunMetadata(t *testing.T) {
	err := testAgent.run("meta-data", testContext)
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if testResource.lastRun != "metadata" {
		t.Fatal("action metadata was not invoked")
	}
}

func TestAgentRunUnknown(t *testing.T) {
	err := testAgent.run("unknown", testContext)
	if err != ErrUnimplemented {
		t.Fatalf("expected ocf.Error (ErrUnimplemented), got %T", err)
	}
}

func BenchmarkAgentRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testAgent.run("start", testContext)
	}
}

func TestAgentRunInvalidArgs(t *testing.T) {
	err := testAgent.Run([]string{}, Env{})
	if err != ErrArgs {
		t.Fatal("expected ocf.Error (ErrArgs), got %T", err)
	}
}

func TestAgentRunInvalidResourceInstance(t *testing.T) {
	err := testAgent.Run([]string{"agent", "start"}, Env{"OCF_RESOURCE_INSTANCE=foobar:NOPE"})
	if _, ok := err.(*strconv.NumError); !ok {
		t.Fatalf("expected *strconv.NumError, got %T", err)
	}
}

func TestAgentRunUnknownArg(t *testing.T) {
	err := testAgent.Run([]string{"agent", "unknown"}, Env{})
	if err != ErrUnimplemented {
		t.Fatalf("expected ocf.Error (ErrUnimplemented), got %T", err)
	}
}

func TestAgentRun(t *testing.T) {
	err := testAgent.Run([]string{"agent", "start"}, testEnv)
	if err != nil {
		t.Fatal("unexpected error:", err)
	} else if testResource.lastRun != "start" {
		t.Fatal("action start was not invoked")
	}
}

func BenchmarkAgentRunCLI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testAgent.Run([]string{"agent", "start"}, testEnv)
	}
}

func TestAgentExitErrorOCF(t *testing.T) {
	w := bytes.NewBuffer(nil)
	log.SetOutput(w)
	log.SetFlags(0)
	defer func() { log.SetOutput(os.Stderr); log.SetFlags(log.LstdFlags) }()

	exitFunc = func(code int) {
		if code != 3 {
			t.Fatal("expected exit code 3, got", code)
		}
	}

	Exit(ErrUnimplemented)

	if s := w.String(); s != "error: action is not implemented\n" {
		t.Fatalf("expected 'error: action is not implemented\n', got '%s'", s)
	}
}

func TestAgentExitError(t *testing.T) {
	w := bytes.NewBuffer(nil)
	log.SetOutput(w)
	log.SetFlags(0)
	defer func() { log.SetOutput(os.Stderr); log.SetFlags(log.LstdFlags) }()

	exitFunc = func(code int) {
		if code != 1 {
			t.Fatal("expected exit code 1, got", code)
		}
	}

	Exit(fmt.Errorf("error message"))

	if s := w.String(); s != "error: error message\n" {
		t.Fatalf("expected 'error: error message\n', got '%s'", s)
	}
}
