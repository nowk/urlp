# urlp

[![Build Status](https://travis-ci.org/nowk/urlp.svg?branch=master)](https://travis-ci.org/nowk/urlp)
[![GoDoc](https://godoc.org/github.com/nowk/urlp?status.svg)](http://godoc.org/github.com/nowk/urlp)

URL pattern match


## Install

    go get gopkg.in/nowk/urlp.v1


## Examples

    v, ok := urlp.Match("/posts/:post_id/comments/:id", "/posts/123/comments/456")
    if !ok {
      // handle
    }

    post_id := v.Get(":post_id")
    id := v.Get(":id")

---

##### .:format

Adding `.:format` at the end of the pattern returns a special named param `:_format` which returns the given format.

    v, ok := urlp.match("/posts/:id.:format", "/posts/123.json")
    if !ok {
      // handle
    }

    id := v.Get(":id")
    format := v.Get(":_format")

With a `.:format` set the pattern will match paths with or without the extension.

    v, ok := urlp.match("/posts/:id.:format", "/posts/123.json")
    // ok       => true
    // :_format => "json"

    v, ok := urlp.match("/posts/:id.:format", "/posts/123")
    // ok       => true
    // :_format => ""

*`.:format` must be at the end of the pattern, else it will be treated as any other named parameter.*

*Without `.:format` paths with a `.ext` will not be matched*


## License

MIT

