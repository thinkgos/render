// Copyright 202 thinkgos.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !noprotopack

package render

import (
	"net/http"

	"github.com/thinkgos/render/render"
)

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func ProtoBuf(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.ProtoBuf{Data: obj})
}
