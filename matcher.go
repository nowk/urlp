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

// Match checks the pattern against the given path, returning any named params
// in the process
func Match(p *Path, s string) (params, bool) {
	l := len(s)
	m := l - 1

	// trim trailing slash
	if l > 1 && s[m] == '/' {
		s = s[:m]

		// decrement counts
		l = l - 1
		m = m - 1
	}

	if s == p.Pattern || (s == "" && p.Pattern == "/") {
		return nil, true
	}

	var ok bool
	var pr params

	d := len(p.Dirs)
	n := 0 // node index

	for i := 0; ; {
		i++
		if i > m {
			break // counted past length of s
		}

		if ok = (n < d); !ok {
			break // has more nodes than available
		}

		j := i
		for ; j < l; j++ {
			if s[j] == '/' {
				break
			}
		}

		a := p.Dirs[n]
		b := s[i-1 : j]

		if a[0] == ':' {
			// lazy alloc
			if pr == nil {
				pr = make(params, 0, p.NoOfParams)
			}

			pr = append(pr, a, b[1:])

		} else {
			if ok = (a == b); !ok {
				break
			}
		}

		i = j // shift i for length
		n++
	}

	if n != d {
		pr = nil // pr[:0]
		ok = false
	}

	return pr, ok
}
