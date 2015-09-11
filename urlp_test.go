package urlp

import (
	"reflect"
	"testing"
)

func TestPattern(t *testing.T) {
	for _, v := range []struct {
		giv string
		exp []node

		noOfParams int
	}{
		{"", []node{"/"}, 0},
		{"/", []node{"/"}, 0},
		{"/f", []node{"/f"}, 0},
		{"/foo", []node{"/foo"}, 0},
		{"/foo/", []node{"/foo"}, 0},
		{"/foo/bar", []node{"/foo", "/bar"}, 0},
		{"/foo/bar/", []node{"/foo", "/bar"}, 0},
		{"/foo/:bar", []node{"/foo", ":bar"}, 2},
		{"/foo/:bar/", []node{"/foo", ":bar"}, 2},
		{"/foo/:bar/baz/:qux", []node{"/foo", ":bar", "/baz", ":qux"}, 4},
		{"/foo/:bar/baz/:qux/", []node{"/foo", ":bar", "/baz", ":qux"}, 4},
		{"/:foo", []node{":foo"}, 2},
		{"/:foo/", []node{":foo"}, 2},
		{"/:foo/*", []node{":foo", "*"}, 2},
	} {
		pat := NewPattern(v.giv)

		if got := pat.Nodes; !reflect.DeepEqual(v.exp, got) {
			t.Errorf("expected %s, got %s", v.exp, got)
		}

		if v.noOfParams != pat.NoOfParams {
			t.Errorf("expected %d, got %d", v.noOfParams, pat.NoOfParams)
		}
	}
}

func TestMatchStaticPaths(t *testing.T) {
	for _, v := range [][]string{
		{"/", ""},
		{"/", "/"},
		{"/posts", "/posts"},
		{"/posts/new", "/posts/new"},
	} {
		for _, ts := range []string{
			"",
			"/",
		} {
			{
				var path = v[1] + ts

				params, ok := NewPattern(v[0]).Match(path)
				if !ok {
					t.Errorf("expected %s to match %s", v[0], path)
				}

				if params != nil {
					t.Errorf("expected no params, got %s", params)
				}
			}
		}
	}
}

func TestMatchDynamicPaths(t *testing.T) {
	for _, v := range [][]string{
		{"/:id", "/123"},
		{"/:id/comments", "/123/comments"},
		{"/posts/:id", "/posts/123"},
		{"/posts/:post_id/comments", "/posts/123/comments"},
		{"/posts/:post_id/comments/:id", "/posts/123/comments/456"},
	} {
		for _, ts := range []string{
			"",
			"/",
		} {
			var path = v[1] + ts

			_, ok := NewPattern(v[0]).Match(path)
			if !ok {
				t.Errorf("expected %s to match %s", v[0], path)
			}
		}
	}
}

func TestParams(t *testing.T) {
	params, ok := NewPattern("/posts/:post_id/comments/:id").
		Match("/posts/123/comments/456")

	var exp = Params{":post_id", "123", ":id", "456"}

	if !ok {
		t.Errorf("expected a match")
	}

	if !reflect.DeepEqual(exp, params) {
		t.Errorf("expected %s, got %s", exp, params)
	}
}

func TestDoesNotMatch(t *testing.T) {
	for _, v := range []struct {
		pat, path string
	}{
		{"/", "/posts"},
		{"/p", "/posts"},
		{"/po", "/posts"},
		{"/pos", "/posts"},
		{"/post", "/posts"},

		{"/posts", "/post"},
		{"/posts", "/pos"},
		{"/posts", "/po"},
		{"/posts", "/p"},
		{"/posts", "/"},

		{"/posts", "/comments"},
		{"/posts", "/posts/123"},

		{"/posts/:post_id", "/posts/123/comments"},
		{"/posts/:post_id", "/posts/123/comments/"},

		{"/posts/:post_id/comments", "/posts/123"},
		{"/posts/:post_id/comments", "/posts/123/"},

		{"/api/label-sets/:label_set_id/labels", "/api/labelsets"},
	} {
		_, ok := NewPattern(v.pat).Match(v.path)
		if ok {
			t.Errorf("expected %s not to match %s", v.pat, v.path)
		}
	}
}

// TestOutOfBounds tests for a panic
func TestOutOfBounds(t *testing.T) {
	_, ok := NewPattern("/posts/:id").Match("/posts")
	if ok {
		t.Errorf("expected no match")
	}
}

func TestMatchWildcard(t *testing.T) {
	for _, v := range []string{
		"/posts/123/comments",
		"/posts/123/comments/456",
		"/posts/123/comments/456/author",
	} {
		_, ok := NewPattern("/posts/:post_id/comments/*").Match(v)
		if !ok {
			t.Errorf("expected a match for %s", v)
		}
	}
}
