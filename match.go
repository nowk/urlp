package urlp

import (
	"strings"
)

type params map[string]string

func (p params) Get(k string) string {
	return p[k]
}

// splits by / removing an starting or ending /
func splits(s string) []string {
	if s == "" || s == "/" {
		return []string{"/"}
	}

	a := 0
	z := len(s)
	if strings.HasPrefix(s, "/") {
		a = 1
	}

	if strings.HasSuffix(s, "/") {
		z = z - 1
	}

	return strings.Split(s[a:z], "/")
}

// isParam checks if a string is a param, which starts with :, eg :post_id
func isParam(s string) bool {
	return strings.HasPrefix(s, ":")
}

// Match matches a pattern to a path string. If the pattern contains param keys,
// eg. :post_id, it will return those associated key:values as a
// params `map[string]string`
func Match(pat, pathStr string) (params, bool) {
	a := splits(pat)
	b := splits(pathStr)
	if len(a) != len(b) {
		return nil, false
	}

	pr := make(map[string]string)

	for i, v := range a {
		n := b[i]
		if isParam(v) { // param
			pr[v[1:]] = n
			continue
		}
		if n != v {
			return nil, false
		}
	}

	if len(pr) == 0 {
		return nil, true
	}

	return pr, true
}
