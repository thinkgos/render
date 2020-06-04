// Copyright (c) 2020 thinkgos<thinkgo@aliyun.com>.All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !noyamlpack

package render

import (
	"net/http"

	"github.com/thinkgos/render/render"
)

// YAML serializes the given struct as YAML into the response body.
func YAML(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.YAML{Data: obj})
}
