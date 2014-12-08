package urlp

import (
	"net/url"
	"strings"
)

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

// Match matches a pattern to a path string. If the pattern contains named
// params, those key:value pairs will be returned as a url.Values
func Match(pat, pathStr string) (url.Values, bool) {
	a := splits(pat)
	b := splits(pathStr)
	if len(a) != len(b) {
		return nil, false
	}

	p := url.Values{}

	for i, v := range a {
		n := b[i]
		if isParam(v) {
			p.Set(v[1:], n)
			continue
		}
		if n != v {
			return nil, false
		}
	}

	if len(p) == 0 {
		return nil, true
	}

	return p, true
}
