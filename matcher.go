package urlp

type params []string

func (p params) Get(k string) string {
	r := len(p) - 2

	for i, v := range p {
		if k == v && i%2 == 0 {
			if i > r {
				return ""
			}

			return p[i+1]
		}
	}

	return ""
}

// dir the first directory level in the path given
func dir(b string) (string, int) {
	for i, v := range b {
		if v == '/' {
			return b[:i], i
		}
	}

	return b, len(b)
}

// trimrs trims trailing slash
func trimrs(s string) (string, int) {
	n := len(s)
	l := n - 1
	if n > 0 && s[l] == '/' {
		return s[:l], l
	}

	return s, n
}

var (
	fk = ".:format"
	fl = 8
)

// formatp parses the pattern and path for `.:format`
func formatp(pattern, path string) (string, string, string, bool) {
	var s int
	path, s = trimrs(path)

	if p := len(pattern); p > fl {
		n := p - fl
		if pattern[n:] == fk {
			pattern = pattern[:n]

			i := s - 1
			for ; i > 0; i-- {
				c := path[i]
				if c == '/' {
					break
				}

				if c == '.' {
					return pattern, path[:i], path[i+1:], true
				}
			}
		}
	}

	return pattern, path, "", false
}

// Match checks the pattern against the given path, returning any named params
// in the process
func Match(pattern, path string) (params, bool) {
	if (path == "" || path == "/") && pattern == "/" {
		return nil, true
	}

	var pr params
	var ok bool
	var f string
	pattern, path, f, ok = formatp(pattern, path)
	if ok {
		pr = append(pr, ":_format", f)
	}
	p := len(pattern)
	s := len(path)

	p_1 := p - 1
	s_1 := s - 1
	var x, y int = 0, 0
	for {
		if x == p && y == s {
			break // when done reaching the end of both paths
		}

		if x > p_1 || y > s_1 {
			return nil, false // if one path has a different number of directory trees
		}

		c := pattern[x]
		if c == ':' {
			k, m := dir(pattern[x:])
			v, n := dir(path[y:])

			x = x + m
			y = y + n

			pr = append(pr, k, v)
			continue
		}

		if c != path[y] {
			return nil, false // if the current chars do nto match
		}

		x++
		y++
	}

	return pr, true
}
