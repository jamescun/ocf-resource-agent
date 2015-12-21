package ocf

import (
	"fmt"
)

type Error struct {
	Exit    int
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("exit status %d: %s", e.Exit, e.Message)
}

// http://www.linux-ha.org/doc/dev-guides/_return_codes.html
var (
	ErrGeneric       = Error{Exit: 1, Message: "unspecified error"}
	ErrArgs          = Error{Exit: 2, Message: "incorrect arguments"}
	ErrUnimplemented = Error{Exit: 3, Message: "action is not implemented"}
	ErrPerm          = Error{Exit: 4, Message: "insufficient permissions"}
	ErrInstalled     = Error{Exit: 5, Message: "required component is missing"}
	ErrConfigured    = Error{Exit: 6, Message: "misconfigured resource"}
	ErrNotRunning    = Error{Exit: 7, Message: "resource not found or not running"}

	ErrRunningMaster = Error{Exit: 8, Message: "resource is master"}
	ErrFailedMaster  = Error{Exit: 9, Message: "resource is not master"}
)
