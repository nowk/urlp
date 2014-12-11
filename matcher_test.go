package urlp

import (
	"github.com/nowk/assert"
	"testing"
)

func TestMatchesIgnoreTrailingSlash(t *testing.T) {
	for _, v := range [][]string{
		{"/", "/"},
		{"/posts", "/posts"},
		{"/posts/new", "/posts/new"},
	} {
		pat, path := v[0], v[1]

		{
			p, ok := Match(pat, path)
			assert.True(t, ok)
			assert.Nil(t, p)
		}
		{
			p, ok := Match(pat, path+"/")
			assert.True(t, ok)
			assert.Nil(t, p)
		}
	}
}

func TestMatchesNamedParamsReturnsParams(t *testing.T) {
	for _, v := range []struct {
		pat, path string
		params    []string
	}{
		{"/posts/:id", "/posts/123", []string{":id", "123"}},
		{"/posts/:post_id/comments/:id", "/posts/123/comments/456", []string{":post_id", "123", ":id", "456"}},
	} {
		{
			p, ok := Match(v.pat, v.path)
			assert.True(t, ok)
			assert.Equal(t, params(v.params), p)
		}
		{
			p, ok := Match(v.pat, v.path+"/")
			assert.True(t, ok)
			assert.Equal(t, params(v.params), p)
		}
	}

}

func TestPathDoesNotMatch(t *testing.T) {
	for _, v := range []struct {
		p, u string
	}{
		{"/", "/posts"},

		{"/posts", "/"},
		{"/posts", "/comments"},
		{"/posts", "/posts/123"},

		{"/posts/:post_id", "/posts/123/comments"},
		{"/posts/:post_id", "/posts/123/comments/"},

		{"/posts/:post_id/comments", "/posts/123"},
		{"/posts/:post_id/comments", "/posts/123/"},
	} {
		p, ok := Match(v.p, v.u)
		assert.False(t, ok, v.p, " ", v.u)
		assert.Nil(t, p)
	}
}

func TestRoot(t *testing.T) {
	for _, v := range []struct {
		p, u string
	}{
		{"/", "/"},
		{"/", ""},
	} {
		p, ok := Match(v.p, v.u)
		assert.True(t, ok, v.p, " != ", v.u)
		assert.Nil(t, p)
	}
}
