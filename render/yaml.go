// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !noyamlpack

package render

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// YAML contains the given interface object.
type YAML struct {
	Data interface{}
}

var _ Render = (*YAML)(nil)
var yamlContentType = []string{"application/x-yaml; charset=utf-8"}

// Render (YAML) marshals the given interface object and writes data with custom ContentType.
func (r YAML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	return yaml.NewEncoder(w).Encode(r.Data)
}

// WriteContentType (YAML) writes YAML ContentType for response.
func (r YAML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, yamlContentType)
}
