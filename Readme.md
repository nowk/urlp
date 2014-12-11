# urlp

[![Build Status](https://travis-ci.org/nowk/urlp.svg?branch=master)](https://travis-ci.org/nowk/urlp)
[![GoDoc](https://godoc.org/github.com/nowk/urlp?status.svg)](http://godoc.org/github.com/nowk/urlp)

URL pattern match


## Examples

    v, ok := urlp.Match("/posts/:post_id/comments/:id", "/posts/123/comments/456")
    if !ok {
      // handle
    }

    post_id := v.Get(":post_id")
    id := v.Get(":id")


## License

MIT

