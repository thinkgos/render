# render this is gin render package, but with option tags

[![GoDoc](https://godoc.org/github.com/things-go/render?status.svg)](https://godoc.org/github.com/things-go/render)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/things-go/render?tab=doc)
[![Build Status](https://www.travis-ci.com/things-go/render.svg?branch=master)](https://www.travis-ci.com/things-go/render)
[![codecov](https://codecov.io/gh/things-go/render/branch/master/graph/badge.svg)](https://codecov.io/gh/things-go/render)
![Action Status](https://github.com/things-go/render/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/things-go/render)](https://goreportcard.com/report/github.com/things-go/render)
[![Licence](https://img.shields.io/github/license/things-go/render)](https://raw.githubusercontent.com/things-go/render/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/things-go/render)](https://github.com/things-go/render/tags)

## Build with [jsoniter](https://github.com/json-iterator/go)

render uses `encoding/json` as default json package, but you can change to [jsoniter](https://github.com/json-iterator/go) by build from other tags.

```sh
$ go build -tags=jsoniter .
```

## Build without 
- [msg](github.com/ugorji/go)
- [protobuf](github.com/golang/protobuf/proto)
- [yaml](https://github.com/go-yaml/yaml)
   
render uses `msg`,`protobuf`,`yaml`, you can build without them with tags.
   
```sh
   $ go build -tags=noprotopack,noyamlpack,nomsgpack .
```

## References
- [gin](https://github.com/gin-gonic/gin)