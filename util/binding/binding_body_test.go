package binding

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindingBody(t *testing.T) {
	for _, tt := range []struct {
		name    string
		binding BindingBody
		body    string
		want    string
	}{
		{
			name:    "ProtoBuf binding",
			binding: ProtoBuf,
			body:    `{"foo":"FOO"}`,
		},
	} {
		t.Logf("testing: %s", tt.name)
		req := requestWithBody("POST", "/", tt.body)
		form := FooStruct{}
		body, _ := ioutil.ReadAll(req.Body)
		assert.NoError(t, tt.binding.BindBody(body, &form))
		assert.Equal(t, FooStruct{"FOO"}, form)
	}
}
