package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"bugangongwei/HelloWorld/util/binding"

	"github.com/gin-gonic/gin"
)

// Bind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//     "application/json" --> JSON binding
//     "application/xml"  --> XML binding
// otherwise --> returns an error.
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// It writes a 400 error and sets Content-Type header "text/plain" in the response if input is not valid.
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return MustBindWith(c, obj, b)
}

// BindQuery is a shortcut for MustBindWith(c, obj, binding.Query), from query.
func BindQuery(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, binding.Query)
}

// MustBindJSONPb is a shortcut for MustBindWith(c, obj, binding.Query), from body.
func MustBindJSONPb(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, binding.ProtoBuf)
}

// ShouldBindJSONPb is a shortcut for ShouldBindWith(c, obj, binding.Query), from body.
func ShouldBindJSONPb(c *gin.Context, obj interface{}) error {
	err := ShouldBindWith(c, obj, binding.ProtoBuf)

	if err != nil && strings.Contains(err.Error(), "EOF") {
		err = errors.New("json in request body should not be empty")
	}

	return err
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func MustBindWith(c *gin.Context, obj interface{}, b binding.Binding) error {
	if err := ShouldBindWith(c, obj, b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind) // nolint: errcheck
		return err
	}
	return nil
}

// ShouldBind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//     "application/json" --> JSON binding
//     "application/xml"  --> XML binding
// otherwise --> returns an error
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// Like c.Bind() but this method does not set the response status code to 400 and abort if the json is not valid.
func ShouldBind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
func ShouldBindQuery(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, binding.Query)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func ShouldBindWith(c *gin.Context, obj interface{}, b binding.Binding) error {
	return b.Bind(c.Request, obj)
}

// ShouldBindBodyWith is similar with ShouldBindWith, but it stores the request
// body into the context, and reuse when it is called again.
//
// NOTE: This method reads the body before binding. So you should use
// ShouldBindWith for better performance if you need to call only once.
func ShouldBindBodyWith(c *gin.Context, obj interface{}, bb binding.BindingBody) (err error) {
	var body []byte
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil {
		body, err = ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		c.Set(gin.BodyBytesKey, body)
	}
	return bb.BindBody(body, obj)
}
