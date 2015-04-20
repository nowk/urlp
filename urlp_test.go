package urlp

import (
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestNewPath(t *testing.T) {
	for _, v := range []struct {
		path   string
		nodes  []node
		params int
	}{
		{"", []node{"/"}, 0},
		{"/", []node{"/"}, 0},
		{"/f", []node{"/f"}, 0},
		{"/foo", []node{"/foo"}, 0},
		{"/foo/", []node{"/foo"}, 0},
		{"/foo/bar", []node{"/foo", "/bar"}, 0},
		{"/foo/bar/", []node{"/foo", "/bar"}, 0},
		{"/foo/:bar", []node{"/foo", ":bar"}, 1},
		{"/foo/:bar/", []node{"/foo", ":bar"}, 1},
		{"/foo/:bar/baz/:qux", []node{"/foo", ":bar", "/baz", ":qux"}, 2},
		{"/foo/:bar/baz/:qux/", []node{"/foo", ":bar", "/baz", ":qux"}, 2},
		{"/:foo", []node{":foo"}, 1},
		{"/:foo/", []node{":foo"}, 1},
	} {
		n := NewPath(v.path)
		assert.Equal(t, v.nodes, n.Nodes, v.path)
		assert.Equal(t, v.params, n.Params)
	}
}

func BenchmarkNewPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewPath("/foo/:bar/baz/:qux/")
	}
}
