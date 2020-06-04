// Package render
package render

import (
	"html/template"
	"io"
	"net/http"

	"github.com/thinkgos/render/render"
)

func Render(w http.ResponseWriter, req *http.Request, code int, r render.Render) {
	w.WriteHeader(code)
	if !bodyAllowedForStatus(code) {
		r.WriteContentType(w)
		return
	}
	if err := r.Render(w); err != nil {
		panic(err)
	}
}

// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}

// HTML renders the HTTP template specified by its file name.
// It also updates the HTTP code and sets the Content-Type as "text/html".
// See http://golang.org/doc/articles/wiki/
func HTML(w http.ResponseWriter, r *http.Request, code int, name string, tpl *template.Template, obj interface{}) {
	Render(w, r, code, render.HTML{
		Template: tpl,
		Name:     name,
		Data:     obj,
	})
}

// IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
// It also sets the Content-Type as "application/json".
// WARNING: we recommend to use this only for development purposes since printing pretty JSON is
// more CPU and bandwidth consuming. Use Context.JSON() instead.
func IndentedJSON(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	Render(w, r, code, render.IndentedJSON{Data: obj})
}

// SecureJSON serializes the given struct as Secure JSON into the response body.
// Default prepends "while(1)," to response body if the given struct is array values.
// It also sets the Content-Type as "application/json".
func SecureJSON(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	Render(w, r, code, render.SecureJSON{Prefix: "while(1);", Data: obj})
}

// JSONP serializes the given struct as JSON into the response body.
// It add padding to response body to request data from a server residing in a different domain than the client.
// It also sets the Content-Type as "application/javascript".
func JSONP(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	callback := r.URL.Query().Get("callback")
	if callback == "" {
		Render(w, r, code, render.JSON{Data: obj})
		return
	}
	Render(w, r, code, render.JsonpJSON{Callback: callback, Data: obj})
}

// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func JSON(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	Render(w, r, code, render.JSON{Data: obj})
}

// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
func AsciiJSON(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	Render(w, r, code, render.AsciiJSON{Data: obj})
}

// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
func PureJSON(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	Render(w, r, code, render.PureJSON{Data: obj})
}

// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
func XML(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	Render(w, r, code, render.XML{Data: obj})
}

// YAML serializes the given struct as YAML into the response body.
func YAML(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	Render(w, r, code, render.YAML{Data: obj})
}

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func ProtoBuf(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	Render(w, r, code, render.ProtoBuf{Data: obj})
}

// String writes the given string into the response body.
func String(w http.ResponseWriter, r *http.Request, code int, format string, values ...interface{}) {
	Render(w, r, code, render.String{Format: format, Data: values})
}

// Redirect returns a HTTP redirect to the specific location.
func Redirect(w http.ResponseWriter, r *http.Request, code int, location string) {
	Render(w, r, http.StatusOK, render.Redirect{
		Code:     code,
		Location: location,
		Request:  r,
	})
}

// Data writes some data into the body stream and updates the HTTP code.
func Data(w http.ResponseWriter, r *http.Request, code int, contentType string, data []byte) {
	Render(w, r, code, render.Data{
		ContentType: contentType,
		Data:        data,
	})
}

// DataFromReader writes the specified reader into the body stream and updates the HTTP code.
func DataFromReader(w http.ResponseWriter, r *http.Request, code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) {
	Render(w, r, code, render.Reader{
		Headers:       extraHeaders,
		ContentType:   contentType,
		ContentLength: contentLength,
		Reader:        reader,
	})
}
