// Copyright (c) 2020 thinkgos<thinkgo@aliyun.com>.All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !nomsgpack

package render

import (
	"net/http"

	"github.com/things-go/render/render"
)

// MsgPack serializes the given struct as Msgpack into the response body.
func MsgPack(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.MsgPack{obj})
}
