package urlp

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
	pat []byte
}

func NewMatcher(pat string) Matcher {
	if pat == "" {
		pat = "/"
	}

	return &matcher{
		pat: []byte(pat),
	}
}

func dir(b []byte) ([]byte, int) {
	for i, v := range b {
		if v == '/' {
			return b[:i], i
		}
	}

	return b, len(b)
}

// Match checks the pattern against the given path, returning any named params
// in the process
func (m *matcher) Match(pathStr string) (params, bool) {
	if pathStr == "" || pathStr == "/" {
		if string(m.pat) == "/" {
			return nil, true
		}
	}

	p, y := m.pat, 0
	u, x := []byte(pathStr), 0

	// trim trailing slash
	n := len(u)
	if u[n-1] == '/' {
		u = u[:n-1]
	}

	var pr params

	for {
		// exit if we make to the end of both paths without issue
		if y == len(p) && x == len(u) {
			break
		}

		if y > len(p)-1 || x > len(u)-1 {
			return nil, false // uneven node lengths
		}

		if p[y] == ':' {
			z, n := dir(p[y:])
			a, m := dir(u[x:])

			y = y + n
			x = x + m

			pr = append(pr, string(z), string(a))
			continue
		}

		if p[y] != u[x] {
			return nil, false // current chars don't match
		}

		if p[y] == u[x] {
			y++
			x++
		}
	}

	return pr, true
}
