# urlp

[![Build Status](https://travis-ci.org/nowk/urlp.svg?branch=master)](https://travis-ci.org/nowk/urlp)
[![GoDoc](https://godoc.org/gopkg.in/nowk/urlp.v3?status.svg)](http://godoc.org/gopkg.in/nowk/urlp.v3)

URL pattern match


## Install

    go get gopkg.in/nowk/urlp.v3


## Examples

    p := urlp.NewPattern("/posts/:post_id/comments/:id")

    v, ok := p.Match("/posts/123/comments/456")
    if !ok {
      // handle
    }

    v.Get(":post_id") // "123"
    v.Get(":id")      // "456"


## License

MIT

