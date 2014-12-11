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

// Match checks the pattern against the given path, returning any named params
// in the process
func Match(pat, pathStr string) (params, bool) {
	if (pathStr == "" || pathStr == "/") && pat == "/" {
		return nil, true
	}

	p, x := pat, 0
	s, y := pathStr, 0

	// trim trailing slash
	n := len(s)
	if s[n-1] == '/' {
		s = s[:n-1]
	}

	var pr params

	plen := len(p)
	slen := len(s)
	plen_1 := plen - 1
	slen_1 := slen - 1

	for {
		if x == plen && y == slen {
			break // when done reaching the end of both paths
		}

		if x > plen_1 || y > slen_1 {
			return nil, false // if one path has a different number of directory trees
		}

		if p[x] == ':' {
			k, m := dir(p[x:])
			v, n := dir(s[y:])

			x = x + m
			y = y + n

			pr = append(pr, k, v)
			continue
		}

		if p[x] != s[y] {
			return nil, false // if the current chars do nto match
		}

		x++
		y++
	}

	return pr, true
}
