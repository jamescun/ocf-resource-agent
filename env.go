package ocf

import (
	"strings"
)

const envSep string = "="

// map of key=value environment variables
type Env []string

// return the value of key name from environment, or empty string
func (e Env) Get(name string) string {
	for i := 0; i < len(e); i++ {
		s := strings.Index(e[i], envSep)
		if s == -1 {
			continue
		}

		if e[i][:s] == name {
			return strings.TrimSpace(e[i][s+1:])
		}
	}

	return ""
}

// return count of fields matching prefix
func (e Env) CountFields(prefix string) (n int) {
	for i := 0; i < len(e); i++ {
		s := strings.Index(e[i], envSep)
		if s == -1 {
			continue
		}

		if strings.HasPrefix(e[i][:s], prefix) {
			n++
		}
	}

	return
}

// return map of fields matching prefix
func (e Env) Fields(prefix string) map[string]string {
	f := make(map[string]string)

	for i := 0; i < len(e); i++ {
		s := strings.Index(e[i], envSep)
		if s == -1 {
			continue
		}

		if strings.HasPrefix(e[i][:s], prefix) {
			f[e[i][len(prefix):s]] = strings.TrimSpace(e[i][s+1:])
		}
	}

	return f
}
