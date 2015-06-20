package urlp

type Path struct {
	Path   string
	Dirs   []string
	Static bool

	// NoOfParams is the # of params multiplied by 2 to represent the k : v pair
	// count
	NoOfParams int
}

func NewPath(s string) *Path {
	d, n := parsePath(s)

	return &Path{
		Path:   s,
		Dirs:   d,
		Static: n == 0,

		NoOfParams: n * 2,
	}
}

func parsePath(s string) ([]string, int) {
	if s == "/" || s == "" {
		return []string{"/"}, 0
	}

	var a []string

	slen := len(s)
	p := 0 // param count
	c := 0 // cursor

	i := 0
	for {
		i++

		if i == slen || s[i] == '/' {
			d := s[c:i]
			if d == "/" {
				break // at last trailing slash
			}
			if d[1] == ':' {
				d = d[1:] // strip off / prefix
				p++       // increment param count
			}

			a = append(a, d)

			c = i // move cursor
		}

		if i >= slen {
			break
		}
	}

	return a, p
}
