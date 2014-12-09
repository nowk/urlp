package urlp

import (
	"strings"
)

type params []string

func (p params) Get(k string) string {
	for i, v := range p {
		if k == v && i%2 == 0 {
			return p[i+1]
		}
	}

	return ""
}

type Matcher interface {
	Match(string) (params, bool)
}

// matcher contains a preconditioned split of the path pattern
type matcher struct {
	pat   string
	split []string
}

func NewMatcher(pat string) Matcher {
	return &matcher{
		pat:   pat,
		split: splits(pat),
	}
}

// splits by / removing an starting or ending /
func splits(s string) []string {
	if s == "" || s == "/" {
		return []string{"/"}
	}

	a := 0
	z := len(s)
	if strings.HasPrefix(s, "/") {
		a = 1
	}

	if strings.HasSuffix(s, "/") {
		z = z - 1
	}

	return strings.Split(s[a:z], "/")
}

// Match checks the pattern against the given path, returning any named params
// in the process
func (m *matcher) Match(pathStr string) (params, bool) {
	b := splits(pathStr)
	if len(m.split) != len(b) {
		return nil, false
	}

	p := make(params, 0, len(b))

	for i, v := range m.split {
		n := b[i]
		if v[0:1] == ":" {
			p = append(p, v, n)
			continue
		}
		if n != v {
			return nil, false
		}
	}

	if len(p) == 0 {
		return nil, true
	}

	return p, true
}
