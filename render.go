// Package render
package render

import (
	"html/template"
	"io"
	"net/http"

	"github.com/thinkgos/render/render"
)

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

// Render render
func Render(w http.ResponseWriter, code int, r render.Render) {
	w.WriteHeader(code)
	if !bodyAllowedForStatus(code) {
		r.WriteContentType(w)
		return
	}

	if err := r.Render(w); err != nil {
		panic(err)
	}
}

// HTML renders the HTTP template specified by its file name.
// It also updates the HTTP code and sets the Content-Type as "text/html".
// See http://golang.org/doc/articles/wiki/
func HTML(w http.ResponseWriter, code int, name string, tpl *template.Template, obj interface{}) {
	Render(w, code, render.HTML{Template: tpl, Name: name, Data: obj})
}

// IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
// It also sets the Content-Type as "application/json".
// WARNING: we recommend to use this only for development purposes since printing pretty JSON is
// more CPU and bandwidth consuming. Use Context.JSON() instead.
func IndentedJSON(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.IndentedJSON{Data: obj})
}

// SecureJSON serializes the given struct as Secure JSON into the response body.
// Default prepends "while(1)," to response body if the given struct is array values.
// It also sets the Content-Type as "application/json".
func SecureJSON(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.SecureJSON{Prefix: "while(1);", Data: obj})
}

// JSONP serializes the given struct as JSON into the response body.
// It add padding to response body to request data from a server residing in a different domain than the client.
// It also sets the Content-Type as "application/javascript".
func JSONP(w http.ResponseWriter, r *http.Request, code int, obj interface{}) {
	callback := r.URL.Query().Get("callback")
	if callback == "" {
		Render(w, code, render.JSON{Data: obj})
		return
	}
	Render(w, code, render.JsonpJSON{Callback: callback, Data: obj})
}

// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func JSON(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.JSON{Data: obj})
}

// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
func AsciiJSON(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.AsciiJSON{Data: obj})
}

// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
func PureJSON(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.PureJSON{Data: obj})
}

// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
func XML(w http.ResponseWriter, code int, obj interface{}) {
	Render(w, code, render.XML{Data: obj})
}

// String writes the given string into the response body.
func String(w http.ResponseWriter, code int, format string, values ...interface{}) {
	Render(w, code, render.String{Format: format, Data: values})
}

// Redirect returns a HTTP redirect to the specific location.
func Redirect(w http.ResponseWriter, r *http.Request, code int, location string) {
	Render(w, http.StatusOK, render.Redirect{Code: code, Location: location, Request: r})
}

// Data writes some data into the body stream and updates the HTTP code.
func Data(w http.ResponseWriter, code int, contentType string, data []byte) {
	Render(w, code, render.Data{ContentType: contentType, Data: data})
}

// DataFromReader writes the specified reader into the body stream and updates the HTTP code.
func DataFromReader(w http.ResponseWriter, code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) {
	Render(w, code, render.Reader{
		Headers:       extraHeaders,
		ContentType:   contentType,
		ContentLength: contentLength,
		Reader:        reader,
	})
}
