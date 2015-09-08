package urlp

type Params []string

func (p Params) Get(k string) string {
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

// Match checks the pattern against the given path, returning any named Params
// in the process
func Match(p *Path, s string) (Params, bool) {
	slen := len(s)
	m := slen - 1

	if slen > 1 && s[m] == '/' {
		s = s[:m] // trim trailing slash

		slen--
		m--
	} else if slen < 1 {
		s = "/"
	}

	// static routes must match exactly, no need to continue after this
	if p.Static {
		return nil, s == p.Path
	}

	var ok bool
	var pr Params

	dlen := len(p.Dirs)
	n := 0 // dir index

	for i := 0; ; {
		if i > m {
			break // done loop
		}

		if ok = (s[i] == '/'); !ok {
			break // next cusror start location is not /
		}

		if ok = (n < dlen); !ok {
			break // has more dirs than available
		}

		i++

		a := p.Dirs[n]
		alen := len(a)
		n++

		if a[0] == ':' {
			if pr == nil {
				pr = make(Params, 0, p.NoOfParams)
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

		if ok = (slen >= i); !ok {
			break // dir is longer than string
		}

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
