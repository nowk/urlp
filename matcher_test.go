package urlp

import (
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestMatchesIgnoreTrailingSlash(t *testing.T) {
	for _, v := range [][]string{
		{"/", "/"},
		{"/posts", "/posts"},
		{"/posts/new", "/posts/new"},
	} {
		pat, path := NewPath(v[0]), v[1]

		{
			p, ok := Match(pat, path)
			assert.True(t, ok, path)
			assert.Nil(t, p)
		}
		{
			p, ok := Match(pat, path+"/")
			assert.True(t, ok, path+"/")
			assert.Nil(t, p)
		}
	}
}

func TestMatchesNamedParamsReturnsParams(t *testing.T) {
	for _, v := range []struct {
		pat, path string
		params    []string
	}{
		{
			"/posts/:id", "/posts/123",
			[]string{":id", "123"},
		},
		{
			"/posts/:post_id/comments/:id", "/posts/123/comments/456",
			[]string{":post_id", "123", ":id", "456"},
		},
	} {
		{
			p, ok := Match(NewPath(v.pat), v.path)
			assert.True(t, ok)
			assert.Equal(t, params(v.params), p)
		}
		{
			p, ok := Match(NewPath(v.pat), v.path+"/")
			assert.True(t, ok)
			assert.Equal(t, params(v.params), p)
		}
	}
}

func TestParamsGetDoesNotIndexOutOfRange(t *testing.T) {
	{
		p := params([]string{":post_id", "123", ":id", "456"})
		assert.Equal(t, "123", p.Get(":post_id"))
		assert.Equal(t, "456", p.Get(":id"))
	}
	{
		p := params([]string{":post_id", "123", ":id"})
		assert.Equal(t, "123", p.Get(":post_id"))
		assert.Equal(t, "", p.Get(":id"))
	}
	{
		p := params([]string{":post_id", "123"})
		assert.Equal(t, "123", p.Get(":post_id"))
		assert.Equal(t, "", p.Get(":id"))
	}
	{
		p := params([]string{":post_id"})
		assert.Equal(t, "", p.Get(":post_id"))
		assert.Equal(t, "", p.Get(":id"))
	}
	{
		p := params([]string{})
		assert.Equal(t, "", p.Get(":post_id"))
		assert.Equal(t, "", p.Get(":id"))
	}
}

func TestPathDoesNotMatch(t *testing.T) {
	for _, v := range []struct {
		p, u string
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
	} {
		_, ok := Match(NewPath(v.p), v.u)
		assert.False(t, ok, v.p, " ", v.u)
		// assert.Nil(t, p)
	}
}

func TestRoot(t *testing.T) {
	for _, v := range []struct {
		p, u string
	}{
		{"/", "/"},
		{"/", ""},
	} {
		p, ok := Match(NewPath(v.p), v.u)
		assert.True(t, ok, v.p, " != ", v.u)
		assert.Nil(t, p)
	}
}

// TODO decide if .:format integration is needed
// func TestFormatParsing(t *testing.T) {
// 	for _, v := range [][]string{
// 		{"/posts.:format", "/posts.html", "html"},
// 		{"/posts.:format", "/posts.json", "json"},
// 		{"/posts.:format", "/posts", ""},

// 		{"/posts/:id.:format", "/posts/123.html", "html"},
// 		{"/posts/:id.:format", "/posts/123.json", "json"},
// 		{"/posts/:id.:format", "/posts/123", ""},
// 	} {
// 		pat, path, format := v[0], v[1], v[2]
// 		p, ok := Match(pat, path)
// 		assert.True(t, ok)
// 		assert.Equal(t, format, p.Get(":_format"))
// 	}
// }

// func TestFormatWithFormatNamedParam(t *testing.T) {
// 	p, ok := Match("/posts/:format/:id.:format", "/posts/long-format/123.html")
// 	assert.True(t, ok)
// 	assert.Equal(t, "html", p.Get(":_format"))
// 	assert.Equal(t, "long-format", p.Get(":format"))
// }

func TestDoesNotIndexOutOfRangeWhenMatchingPathToEmptyString(t *testing.T) {
	_, ok := Match(NewPath("/posts"), "")
	assert.False(t, ok)
}
