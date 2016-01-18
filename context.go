package ocf

import (
	"strconv"
	"strings"
)

// resource environment invoked from Pacemaker
type Context struct {
	// OCF resource instance literal
	// http://www.linux-ha.org/doc/dev-guides/_literal_ocf_resource_instance_literal.html
	Resource ResourceInstance

	// keys from OCF_RESKEY_<name> environment variables
	Keys map[string]string

	// notification type (type-operation)
	Notify string
}

func (c *Context) readNotifyType() {
	c.Notify = strings.Join([]string{c.Keys["meta_notify_type"], c.Keys["meta_notify_operation"]}, "-")
}

// return key string
func (c *Context) String(name string) string {
	return c.Keys[name]
}

// return key as array of strings
func (c *Context) StringArray(name string) []string {
	return strings.Fields(c.Keys[name])
}

// return int from key string (0 if invalid)
func (c *Context) Int(name string) int {
	v, _ := strconv.Atoi(c.String(name))
	return v
}

// return int64 from key string (0 if invalid)
func (c *Context) Int64(name string) int64 {
	v, _ := strconv.ParseInt(c.String(name), 10, 64)
	return v
}
