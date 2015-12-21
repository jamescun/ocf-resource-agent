package crm

import (
	"os/exec"
	"strconv"
	"strings"
)

const (
	// http://clusterlabs.org/doc/en-US/Pacemaker/1.1/html/Pacemaker_Explained/ch06.html#_scores
	INFINITY = 1000000
)

type CRM struct {
	// property primitive name for attribute set
	Property string

	// path to crm_master binary
	MasterPath string

	// path to crm_attribute binary
	AttrPath string
}

// return new CRM object for property
func New(property string) *CRM {
	return &CRM{
		Property:   property,
		MasterPath: "crm_master",
		AttrPath:   "crm_attribute",
	}
}

// return value of property key in corosync
func (c CRM) Get(name string) (string, error) {
	cmd := exec.Command(c.AttrPath, "--type", "crm_config", "--name", name, "--set-name", c.Property, "--query", "-q")

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(v)), nil
}

// set the value of property key in corosync
func (c CRM) Set(name, value string) error {
	cmd := exec.Command(c.AttrPath, "--type", "crm_config", "--name", name, "--set-name", c.Property, "--update", value)
	return cmd.Run()
}

// set local node master preference score
func (c CRM) SetScore(score int) error {
	cmd := exec.Command(c.MasterPath, "-l", "reboot", "-v", strconv.Itoa(score))
	return cmd.Run()
}

// get local node master preference score
func (c CRM) GetScore() (int, error) {
	cmd := exec.Command(c.MasterPath, "-l", "reboot", "--get-value", "--quiet")

	v, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(string(v)))
}

// delete local node master preference score
func (c CRM) DeleteScore() error {
	cmd := exec.Command(c.MasterPath, "-l", "reboot", "-D")
	return cmd.Run()
}

var DefaultCRM = New("replication")

func Get(name string) (string, error) { return DefaultCRM.Get(name) }
func Set(name, value string) error    { return DefaultCRM.Set(name, value) }
func SetScore(score int) error        { return DefaultCRM.SetScore(score) }
func GetScore() (int, error)          { return DefaultCRM.GetScore() }
func DeleteScore() error              { return DefaultCRM.DeleteScore() }
