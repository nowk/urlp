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
func trimrs(s *string) {
	n := len(*s)
	if (*s)[n-1] == '/' {
		*s = (*s)[:n-1]
	}
}

// Match checks the pattern against the given path, returning any named params
// in the process
func Match(pattern, path string) (params, bool) {
	if (path == "" || path == "/") && pattern == "/" {
		return nil, true
	}
	trimrs(&path)

	p := len(pattern)
	s := len(path)
	p_1 := p - 1
	s_1 := s - 1

	var pr params
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
