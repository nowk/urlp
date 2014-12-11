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
	if s[l] == '/' {
		return s[:l], l
	}

	return s, n
}

// Match checks the pattern against the given path, returning any named params
// in the process
func Match(pattern, path string) (params, bool) {
	if (path == "" || path == "/") && pattern == "/" {
		return nil, true
	}

	var pr params
	var s int
	path, s = trimrs(path)
	p := len(pattern)

	if p > formatPatlen {
		if f := pattern[p-formatPatlen:]; f == formatPat {
			p = p - formatPatlen
			pattern = pattern[:p]

			i := s - 1
			for {
				if i == 0 {
					break
				}

				c := path[i]
				if c == '.' {
					ext := path[i+1:]
					pr = append(pr, ":_format", ext)

					path = path[:i]
					s = i
				}

				if c == '/' {
					break // if reached directory, no format, exit
				}

				i--
			}
		}
	}

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

		if pattern[x] == ':' {
			k, m := dir(pattern[x:])
			v, n := dir(path[y:])

			x = x + m
			y = y + n

			pr = append(pr, k, v)
			continue
		}

		if pattern[x] != path[y] {
			return nil, false // if the current chars do nto match
		}

		x++
		y++
	}

	return pr, true
}
