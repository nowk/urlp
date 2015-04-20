package urlp

type node string

func (n node) IsParam() bool {
	if len(n) <= 1 {
		return false
	}

	return n[1] == ':'
}

type Path struct {
	Pattern string
	Nodes   []node
	Params  int
}

func NewPath(pat string) *Path {
	nodes, n := parseNodes(pat)

	return &Path{
		Pattern: pat,
		Nodes:   nodes,
		Params:  n,
	}
}

// parseNodes parses string into individual node sections, eg /<path> and
// returns the number of nodes that are param based nodes eg /:<param>.
func parseNodes(pat string) ([]node, int) {
	if pat == "/" || pat == "" {
		return []node{"/"}, 0
	}

	n := make([]node, 0, nodeCount(pat))
	m := len(pat)
	p := 0
	i := 0
	c := 0
	for {
		c++

		if c == m || pat[c] == '/' {
			no := node(pat[i:c])
			if no == "/" {
				break
			}
			if no[1] == ':' {
				no = no[1:] // strip off / prefix
				p++
			}

			n = append(n, no)
			i = c
		}

		if c >= m {
			break
		}
	}

	return n, p
}

// nodeCount returns the number of nodes, eg /<path> sections available in a
// given string. The last slash is not counted as a node.
func nodeCount(pat string) int {
	var lastchar rune
	n := 0
	for _, v := range pat {
		if v == '/' {
			n++
		}

		lastchar = v
	}

	if lastchar == '/' {
		n--
	}

	return n
}
