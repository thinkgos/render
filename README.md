# render this is gin render package, but with option tags

[![GoDoc](https://godoc.org/github.com/thinkgos/render?status.svg)](https://godoc.org/github.com/thinkgos/render)
[![Build Status](https://www.travis-ci.org/thinkgos/render.svg?branch=master)](https://www.travis-ci.org/thinkgos/render)
[![codecov](https://codecov.io/gh/thinkgos/render/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/render)
![Action Status](https://github.com/thinkgos/render/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/render)](https://goreportcard.com/report/github.com/thinkgos/render)
[![Licence](https://img.shields.io/github/license/thinkgos/render)](https://raw.githubusercontent.com/thinkgos/render/master/LICENSE)

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