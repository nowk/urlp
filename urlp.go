package urlp

type node string

func (n node) IsParam() bool {
	if len(n) < 1 {
		return false
	}

	return n[0] == ':'
}

func (n node) IsWildcard() bool {
	return n == "*"
}

type Pattern struct {
	Path       string
	Nodes      []node
	Static     bool
	NoOfParams int
}

func NewPattern(path string) *Pattern {
	if path == "/" || path == "" {
		return &Pattern{
			Path:   path,
			Nodes:  []node{"/"},
			Static: true,
		}
	}

	var (
		nodes []node

		lenPath    = len(path)
		noOfParams = 0
	)

	c := 0 // cursor
	i := 0
	for {
		i++

		if i == lenPath || path[i] == '/' {
			part := path[c:i]

			if part == "/" {
				break // at last trailing slash
			}

			if char := part[1]; char == ':' || char == '*' {
				part = part[1:] // strip off / prefix

				if char == ':' {
					noOfParams++
				}
			}

			nodes = append(nodes, node(part))

			c = i // move cursor
		}

		if i >= lenPath {
			break
		}
	}

	return &Pattern{
		Path:       path,
		Nodes:      nodes,
		Static:     noOfParams == 0,
		NoOfParams: noOfParams * 2,
	}
}

func (p *Pattern) Match(path string) (Params, bool) {
	lenPath := len(path)
	z := lenPath - 1

	// normalize the path for ts
	if lenPath > 1 && path[z] == '/' {
		path = path[:z] // trim ts
		lenPath--
		z--
	} else if path == "" {
		path = "/"
	}

	// static, match exactly
	if p.Static {
		return nil, p.Path == path
	}

	var ok bool = true
	var pr Params

	c := 0 // cursor location

	for _, v := range p.Nodes {
		if v.IsWildcard() {
			c = lenPath
			break
		}
		if v.IsParam() {
			c++
			if ok = !(c > lenPath); !ok {
				break
			}

			i := c
			for {
				if c > z || path[c] == '/' {
					pr = append(pr, string(v), path[i:c])

					break
				}
				c++
			}

			continue
		}

		i := c
		c = c + len(v)

		if ok = !(c > lenPath); !ok {
			break // if cursor + node length exceed the length of the path
		}
		if ok = (string(v) == path[i:c]); !ok {
			break // if the current node does not match the same node in path
		}
	}

	return pr, (c == lenPath)
}

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
