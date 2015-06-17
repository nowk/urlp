package urlp

import (
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestNewPathReturnsAParsedParth(t *testing.T) {
	for _, v := range []struct {
		path       string
		dirs       []string
		noOfParams int
	}{
		{"", []string{"/"}, 0},
		{"/", []string{"/"}, 0},
		{"/f", []string{"/f"}, 0},
		{"/foo", []string{"/foo"}, 0},
		{"/foo/", []string{"/foo"}, 0},
		{"/foo/bar", []string{"/foo", "/bar"}, 0},
		{"/foo/bar/", []string{"/foo", "/bar"}, 0},
		{"/foo/:bar", []string{"/foo", ":bar"}, 2},
		{"/foo/:bar/", []string{"/foo", ":bar"}, 2},
		{"/foo/:bar/baz/:qux", []string{"/foo", ":bar", "/baz", ":qux"}, 4},
		{"/foo/:bar/baz/:qux/", []string{"/foo", ":bar", "/baz", ":qux"}, 4},
		{"/:foo", []string{":foo"}, 2},
		{"/:foo/", []string{":foo"}, 2},
	} {
		n := NewPath(v.path)
		assert.Equal(t, v.dirs, n.Dirs, v.path)
		assert.Equal(t, v.noOfParams, n.NoOfParams)
	}
}
