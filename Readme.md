# urlp

[![Build Status](https://travis-ci.org/nowk/urlp.svg?branch=master)](https://travis-ci.org/nowk/urlp)
[![GoDoc](https://godoc.org/github.com/nowk/urlp?status.svg)](http://godoc.org/github.com/nowk/urlp)

URL pattern match


## Examples

    p := "/posts/:post_id/comments/:id"
    u := "/posts/123/comments/456"

    m := urlp.NewMatcher(p)
    v, ok := m.Match(u)
    if !ok {
      // handle url doesn't match pattern
    }

    post_id := v.Get(":post_id")
    id := v.Get(":id")


## License

MIT

