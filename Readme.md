# urlp

[![Build Status](https://travis-ci.org/nowk/urlp.svg?branch=master)](https://travis-ci.org/nowk/urlp)
[![GoDoc](https://godoc.org/github.com/nowk/urlp?status.svg)](http://godoc.org/github.com/nowk/urlp)

URL pattern match


## Examples

    p := "/posts/:post_id/comments/:id"
    u := "/posts/123/comments/456"

    m, ok := urlp.Matches(p, u)
    if !ok {
      // handle url doesn't match pattern
    }

    post_id := m.Get("post_id")
    id := m.Get("id")


## License

MIT

