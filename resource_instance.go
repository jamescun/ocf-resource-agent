package ocf

import (
	"strconv"
	"strings"
)

// OCF resouce instance literal
// http://www.linux-ha.org/doc/dev-guides/_literal_ocf_resource_instance_literal.html
type ResourceInstance struct {
	Name  string
	Clone int
}

// return ResourceInstance literal from string
func ParseResourceInstance(s string) (ri ResourceInstance, err error) {
	if len(s) == 0 {
		err = ErrUnimplemented
		return
	}

	i := strings.Index(s, ":")
	if i == -1 {
		// no clone identifier
		ri.Name = strings.TrimSpace(s)
		return
	}

	ri.Clone, err = strconv.Atoi(strings.TrimSpace(s[i+1:]))
	if err != nil {
		return
	}

	ri.Name = strings.TrimSpace(s[:i])
	return
}
