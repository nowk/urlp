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
	slen := len(s)
	m := slen - 1

	if slen > 1 && s[m] == '/' {
		s = s[:m] // trim trailing slash

		slen--
		m--
	}

	if s == p.Pattern || (s == "" && p.Pattern == "/") {
		return nil, true
	}

	var ok bool
	var pr params

	dlen := len(p.Dirs)
	n := 0 // dir index

	for i := 0; ; {
		if i > m {
			break // counted past length of s
		}

		if ok = (s[i] == '/'); !ok {
			break // next cusror start location is not /
		}

		i++

		if ok = (n < dlen); !ok {
			break // has more dirs than available
		}

		a := p.Dirs[n]
		alen := len(a)
		n++

		if ok = (slen > alen); !ok {
			break // dir is longer than string
		}

		if a[0] == ':' {
			if pr == nil {
				pr = make(params, 0, p.NoOfParams)
			}

			h := i
			for ; i < slen; i++ {
				if s[i] == '/' {
					break
				}
			}

			pr = append(pr, a, s[h:i])

			continue
		}

		h := i - 1
		i = h + alen

		b := s[h:i]
		if ok = (a == b); !ok {
			break // does not match
		}
	}

	// check to make sure the num of dirs iterated through matches the num of dirs
	// in the Path object
	if n != dlen {
		ok = false
	}

	return pr, ok
}
