package ocf

import (
	"log"
	"os"
	"strings"
)

type Agent struct {
	r Resource
}

// initiate new Agent to represent resource
func New(r Resource) *Agent {
	return &Agent{r: r}
}

// return error from executed resource handler
func (a *Agent) run(command string, ctx *Context) error {
	switch command {
	case "start":
		return a.r.Start(ctx)

	case "stop":
		return a.r.Stop(ctx)

	case "monitor":
		return a.r.Monitor(ctx)

	case "notify":
		return a.r.Notify(ctx)

	case "promote":
		return a.r.Promote(ctx)

	case "demote":
		return a.r.Demote(ctx)

	case "meta-data":
		return a.r.Metadata()
	}

	return ErrUnimplemented
}

// run agent with argv parameters and environment variables.
// returns nil on success
func (a *Agent) Run(args []string, env Env) (err error) {
	if len(args) < 2 {
		err = ErrArgs
		return
	}

	ri, err := ParseResourceInstance(env.Get("OCF_RESOURCE_INSTANCE"))
	if err != nil && err != ErrUnimplemented {
		return
	}

	ctx := &Context{
		Resource: ri,
		Keys:     env.Fields("OCF_RESKEY_"),
	}

	err = a.run(strings.TrimSpace(args[1]), ctx)
	if err != nil {
		return
	}

	return nil
}

// this is a hack to get around terminating while
// executing tests
var exitFunc func(int) = os.Exit

// accepts an Error or runtime.error, printing appropriate error message to STDERR
// and exiting with relevant OCF exit code
func Exit(err error) {
	var status int = 0

	switch t := err.(type) {
	case Error:
		log.Println("error:", t.Message)
		status = t.Exit

	case error:
		log.Println("error:", t.Error())
		status = 1
	}

	exitFunc(status)
}
