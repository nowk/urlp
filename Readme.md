# urlp

[![Build Status](https://travis-ci.org/nowk/urlp.svg?branch=master)](https://travis-ci.org/nowk/urlp)
[![GoDoc](https://godoc.org/gopkg.in/nowk/urlp.v2?status.svg)](http://godoc.org/gopkg.in/nowk/urlp.v2)

URL pattern match


## Install

    go get gopkg.in/nowk/urlp.v2


## Examples

    p := urlp.NewPath("/posts/:post_id/comments/:id")

    v, ok := urlp.Match(p, "/posts/123/comments/456")
    if !ok {
      // handle
    }

    v.Get(":post_id") // "123"
    v.Get(":id")      // "456"


## License

MIT

