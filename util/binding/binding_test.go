// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

type appkey struct {
	Appkey string `json:"appkey" form:"appkey" protobuf:"varint,3,opt,name=appkey,json=appkey,proto3"`
}

type QueryTest struct {
	Page int `json:"page" form:"page" protobuf:"varint,1,opt,name=page,json=page,proto3"`
	Size int `json:"size" form:"size" protobuf:"varint,2,opt,name=size,json=size,proto3"`
	appkey
}

func (fd *QueryTest) Reset()         {}
func (fd *QueryTest) String() string { return "" }
func (fd *QueryTest) ProtoMessage()  {}

type FooStruct struct {
	Foo string `msgpack:"foo" json:"foo" form:"foo" xml:"foo" binding:"required" protobuf:"bytes,1,opt,name=foo,json=foo,proto3"`
}

func (fd *FooStruct) Reset()         {}
func (fd *FooStruct) String() string { return "" }
func (fd *FooStruct) ProtoMessage()  {}

type FooBarStruct struct {
	FooStruct
	Bar string `msgpack:"bar" json:"bar" form:"bar" xml:"bar" binding:"required" protobuf:"varint,2,opt,name=bar,json=bar,proto3"`
}

func (fd *FooBarStruct) Reset()         {}
func (fd *FooBarStruct) String() string { return "" }
func (fd *FooBarStruct) ProtoMessage()  {}

type FooBarFileStruct struct {
	FooBarStruct
	File *multipart.FileHeader `form:"file" binding:"required" protobuf:"varint,3,opt,name=file,json=file,proto3"`
}

func (fd *FooBarFileStruct) Reset()         {}
func (fd *FooBarFileStruct) String() string { return "" }
func (fd *FooBarFileStruct) ProtoMessage()  {}

type FooBarFileFailStruct struct {
	FooBarStruct
	File *multipart.FileHeader `invalid_name:"file" binding:"required"`
}

type FooDefaultBarStruct struct {
	FooStruct
	Bar string `msgpack:"bar" json:"bar" form:"bar,default=hello" xml:"bar" binding:"required" protobuf:"varint,2,opt,name=bar,json=bar,proto3"`
}

func (fd *FooDefaultBarStruct) Reset()         {}
func (fd *FooDefaultBarStruct) String() string { return "" }
func (fd *FooDefaultBarStruct) ProtoMessage()  {}

type FooStructForMapType struct {
	MapFoo map[string]interface{} `form:"map_foo" protobuf:"varint,1,opt,name=map_foo,json=map_foo,proto3"`
}

type FooStructForIgnoreFormTag struct {
	Foo *string `form:"-"`
}

type FooStructForSliceType struct {
	SliceFoo []int `form:"slice_foo" protobuf:"varint,1,opt,name=slice_foo,json=slice_foo,proto3"`
}

type FooStructForStructType struct {
	StructFoo struct {
		Idx int `form:"idx" protobuf:"varint,1,opt,name=idx,json=idx,proto3"`
	}
}

type FooStructForStructPointerType struct {
	StructPointerFoo *struct {
		Name string `form:"name" protobuf:"varint,1,opt,name=name,json=name,proto3"`
	}
}

type FooStructForSliceMapType struct {
	// Unknown type: not support map
	SliceMapFoo []map[string]interface{} `form:"slice_map_foo" protobuf:"varint,1,opt,name=slice_map_foo,json=slice_map_foo,proto3"`
}

type FooStructForBoolType struct {
	BoolFoo bool `form:"bool_foo" protobuf:"varint,1,opt,name=bool_foo,json=bool_foo,proto3"`
}

type FooStructForStringPtrType struct {
	PtrFoo *string `form:"ptr_foo" protobuf:"varint,1,opt,name=ptr_foo,json=ptr_foo,proto3"`
	PtrBar *string `form:"ptr_bar" binding:"required" protobuf:"varint,2,opt,name=ptr_bar,json=ptr_bar,proto3"`
}

type FooStructForMapPtrType struct {
	PtrBar *map[string]interface{} `form:"ptr_bar" protobuf:"varint,1,opt,name=ptr_bar,json=ptr_bar,proto3"`
}

func TestBindingDefault(t *testing.T) {
	assert.Equal(t, Form, Default("GET", ""))
	assert.Equal(t, Form, Default("GET", MIMEJSON))

	assert.Equal(t, ProtoBuf, Default("POST", MIMEJSON))
	assert.Equal(t, ProtoBuf, Default("PUT", MIMEJSON))

	assert.Equal(t, Form, Default("POST", MIMEPOSTForm))
	assert.Equal(t, Form, Default("PUT", MIMEPOSTForm))

	assert.Equal(t, FormMultipart, Default("POST", MIMEMultipartPOSTForm))
	assert.Equal(t, FormMultipart, Default("PUT", MIMEMultipartPOSTForm))

	assert.Equal(t, ProtoBuf, Default("POST", MIMEPROTOBUF))
	assert.Equal(t, ProtoBuf, Default("PUT", MIMEPROTOBUF))
}

func TestBindingJSONNilBody(t *testing.T) {
	var obj FooStruct
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	err := ProtoBuf.Bind(req, &obj)
	assert.Error(t, err)
}

func TestBindingJSON(t *testing.T) {
	testBodyBinding(t,
		ProtoBuf, "protobuf",
		"/", "/",
		`{"foo": "bar"}`, `{"bar": "foo"}`)
}

func TestBindingForm(t *testing.T) {
	testFormBinding(t, "POST",
		"/", "/",
		"foo=bar&bar=foo", "bar2=foo")
}

func TestBindingForm2(t *testing.T) {
	testFormBinding(t, "GET",
		"/?foo=bar&bar=foo", "/?bar2=foo",
		"", "")
}

func TestBindingFormEmbeddedStruct(t *testing.T) {
	testFormBindingEmbeddedStruct(t, "POST",
		"/", "/",
		"page=1&size=2&appkey=test-appkey", "bar2=foo")
}

func TestBindingFormEmbeddedStruct2(t *testing.T) {
	testFormBindingEmbeddedStruct(t, "GET",
		"/?page=1&size=2&appkey=test-appkey", "/?bar2=foo",
		"", "")
}

func TestBindingFormDefaultValue(t *testing.T) {
	testFormBindingDefaultValue(t, "POST",
		"/", "/",
		"foo=bar&bar=hello", "bar2=foo")
}

func TestBindingFormDefaultValue2(t *testing.T) {
	testFormBindingDefaultValue(t, "GET",
		"/?foo=bar&bar=hello", "/?bar2=foo",
		"", "")
}

func TestFormBindingIgnoreField(t *testing.T) {
	testFormBindingIgnoreField(t, "POST",
		"/", "/",
		"-=bar", "")
}

func TestBindingFormForType(t *testing.T) {
	testFormBindingForType(t, "POST",
		"/", "/",
		"map_foo={\"bar\":123}", "map_foo=1", "Map")

	testFormBindingForType(t, "POST",
		"/", "/",
		"slice_foo=1&slice_foo=2", "bar2=1&bar2=2", "Slice")

	testFormBindingForType(t, "GET",
		"/?slice_foo=1&slice_foo=2", "/?bar2=1&bar2=2",
		"", "", "Slice")

	testFormBindingForType(t, "POST",
		"/", "/",
		"slice_map_foo=1&slice_map_foo=2", "bar2=1&bar2=2", "SliceMap")

	testFormBindingForType(t, "GET",
		"/?slice_map_foo=1&slice_map_foo=2", "/?bar2=1&bar2=2",
		"", "", "SliceMap")

	testFormBindingForType(t, "POST",
		"/", "/",
		"ptr_bar=test", "bar2=test", "Ptr")

	testFormBindingForType(t, "GET",
		"/?ptr_bar=test", "/?bar2=test",
		"", "", "Ptr")

	testFormBindingForType(t, "POST",
		"/", "/",
		"idx=123", "id1=1", "Struct")

	testFormBindingForType(t, "GET",
		"/?idx=123", "/?id1=1",
		"", "", "Struct")

	testFormBindingForType(t, "POST",
		"/", "/",
		"name=thinkerou", "name1=ou", "StructPointer")

	testFormBindingForType(t, "GET",
		"/?name=thinkerou", "/?name1=ou",
		"", "", "StructPointer")
}

func TestBindingQuery(t *testing.T) {
	testQueryBinding(t, "POST",
		"/?foo=bar&bar=foo", "/",
		"foo=unused", "bar2=foo")
}

func TestBindingQuery2(t *testing.T) {
	testQueryBinding(t, "GET",
		"/?foo=bar&bar=foo", "/?bar2=foo",
		"foo=unused", "")
}

func TestBindingQueryFail(t *testing.T) {
	testQueryBindingFail(t, "POST",
		"/?map_foo=", "/",
		"map_foo=unused", "bar2=foo")
}

func TestBindingQueryFail2(t *testing.T) {
	testQueryBindingFail(t, "GET",
		"/?map_foo=", "/?bar2=foo",
		"map_foo=unused", "")
}

func TestBindingQueryBoolFail(t *testing.T) {
	testQueryBindingBoolFail(t, "GET",
		"/?bool_foo=fasl", "/?bar2=foo",
		"bool_foo=unused", "")
}

func createFormPostRequest(t *testing.T) *http.Request {
	req, err := http.NewRequest("POST", "/?foo=getfoo&bar=getbar", bytes.NewBufferString("foo=bar&bar=foo"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEPOSTForm)
	return req
}

func createDefaultFormPostRequest(t *testing.T) *http.Request {
	req, err := http.NewRequest("POST", "/?foo=getfoo&bar=getbar", bytes.NewBufferString("foo=bar&bar=hello"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEPOSTForm)
	return req
}

func createFormPostRequestForMap(t *testing.T) *http.Request {
	req, err := http.NewRequest("POST", "/?map_foo=getfoo", bytes.NewBufferString("map_foo={\"bar\":123}"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEPOSTForm)
	return req
}

func createFormPostRequestForMapFail(t *testing.T) *http.Request {
	req, err := http.NewRequest("POST", "/?map_foo=getfoo", bytes.NewBufferString("map_foo=hello"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEPOSTForm)
	return req
}

func createFormFilesMultipartRequest(t *testing.T) *http.Request {
	boundary := "--testboundary"
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	assert.NoError(t, mw.SetBoundary(boundary))
	assert.NoError(t, mw.WriteField("foo", "bar"))
	assert.NoError(t, mw.WriteField("bar", "foo"))

	f, err := os.Open("form.go")
	assert.NoError(t, err)
	defer f.Close()
	fw, err1 := mw.CreateFormFile("file", "form.go")
	assert.NoError(t, err1)
	_, err = io.Copy(fw, f)
	assert.Nil(t, err)

	req, err2 := http.NewRequest("POST", "/?foo=getfoo&bar=getbar", body)
	assert.NoError(t, err2)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+"; boundary="+boundary)

	return req
}

func createFormFilesMultipartRequestFail(t *testing.T) *http.Request {
	boundary := "--testboundary"
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	assert.NoError(t, mw.SetBoundary(boundary))
	assert.NoError(t, mw.WriteField("foo", "bar"))
	assert.NoError(t, mw.WriteField("bar", "foo"))

	f, err := os.Open("form.go")
	assert.NoError(t, err)
	defer f.Close()
	fw, err1 := mw.CreateFormFile("file_foo", "form_foo.go")
	assert.NoError(t, err1)
	_, err = io.Copy(fw, f)
	assert.Nil(t, err)

	req, err2 := http.NewRequest("POST", "/?foo=getfoo&bar=getbar", body)
	assert.NoError(t, err2)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+"; boundary="+boundary)

	return req
}

func createFormMultipartRequest(t *testing.T) *http.Request {
	boundary := "--testboundary"
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	assert.NoError(t, mw.SetBoundary(boundary))
	assert.NoError(t, mw.WriteField("foo", "bar"))
	assert.NoError(t, mw.WriteField("bar", "foo"))
	req, err := http.NewRequest("POST", "/?foo=getfoo&bar=getbar", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+"; boundary="+boundary)
	return req
}

func createFormMultipartRequestForMap(t *testing.T) *http.Request {
	boundary := "--testboundary"
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	assert.NoError(t, mw.SetBoundary(boundary))
	assert.NoError(t, mw.WriteField("map_foo", "{\"bar\":123, \"name\":\"thinkerou\", \"pai\": 3.14}"))
	req, err := http.NewRequest("POST", "/?map_foo=getfoo", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+"; boundary="+boundary)
	return req
}

func createFormMultipartRequestForMapFail(t *testing.T) *http.Request {
	boundary := "--testboundary"
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	assert.NoError(t, mw.SetBoundary(boundary))
	assert.NoError(t, mw.WriteField("map_foo", "3.14"))
	req, err := http.NewRequest("POST", "/?map_foo=getfoo", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+"; boundary="+boundary)
	return req
}

func TestBindingFormPost(t *testing.T) {
	req := createFormPostRequest(t)
	var obj FooBarStruct
	assert.NoError(t, FormPost.Bind(req, &obj))

	assert.Equal(t, "form-urlencoded", FormPost.Name())
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, "foo", obj.Bar)
}

func TestBindingDefaultValueFormPost(t *testing.T) {
	req := createDefaultFormPostRequest(t)
	var obj FooDefaultBarStruct
	assert.NoError(t, FormPost.Bind(req, &obj))

	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, "hello", obj.Bar)
}

func TestBindingFormPostForMap(t *testing.T) {
	req := createFormPostRequestForMap(t)
	var obj FooStructForMapType
	err := FormPost.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, float64(123), obj.MapFoo["bar"].(float64))
}

func TestBindingFormPostForMapFail(t *testing.T) {
	req := createFormPostRequestForMapFail(t)
	var obj FooStructForMapType
	err := FormPost.Bind(req, &obj)
	assert.Error(t, err)
}

func TestBindingFormFilesMultipart(t *testing.T) {
	req := createFormFilesMultipartRequest(t)
	var obj FooBarFileStruct
	err := FormMultipart.Bind(req, &obj)
	assert.Nil(t, err)
	// file from os
	f, _ := os.Open("form.go")
	defer f.Close()
	fileActual, _ := ioutil.ReadAll(f)

	// file from multipart
	mf, _ := obj.File.Open()
	defer mf.Close()
	fileExpect, _ := ioutil.ReadAll(mf)

	assert.Equal(t, FormMultipart.Name(), "multipart/form-data")
	assert.Equal(t, obj.Foo, "bar")
	assert.Equal(t, obj.Bar, "foo")
	assert.Equal(t, fileExpect, fileActual)
}

func TestBindingFormFilesMultipartFail(t *testing.T) {
	req := createFormFilesMultipartRequestFail(t)
	var obj FooBarFileFailStruct
	err := FormMultipart.Bind(req, &obj)
	assert.Error(t, err)
}

func TestBindingFormMultipart(t *testing.T) {
	req := createFormMultipartRequest(t)
	var obj FooBarStruct
	assert.NoError(t, FormMultipart.Bind(req, &obj))

	assert.Equal(t, "multipart/form-data", FormMultipart.Name())
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, "foo", obj.Bar)
}

func TestBindingFormMultipartForMap(t *testing.T) {
	req := createFormMultipartRequestForMap(t)
	var obj FooStructForMapType
	err := FormMultipart.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, float64(123), obj.MapFoo["bar"].(float64))
	assert.Equal(t, "thinkerou", obj.MapFoo["name"].(string))
	assert.Equal(t, float64(3.14), obj.MapFoo["pai"].(float64))
}

func TestBindingFormMultipartForMapFail(t *testing.T) {
	req := createFormMultipartRequestForMapFail(t)
	var obj FooStructForMapType
	err := FormMultipart.Bind(req, &obj)
	assert.Error(t, err)
}

func TestBindingProtoBufFail(t *testing.T) {
	test := &protoexample.Test{
		Label: proto.String("yes"),
	}
	data, _ := proto.Marshal(test)

	testProtoBodyBindingFail(t,
		ProtoBuf, "protobuf",
		"/", "/",
		string(data), string(data[1:]))
}

func TestValidationFails(t *testing.T) {
	var obj FooStruct
	req := requestWithBody("POST", "/", `{"bar": "foo"}`)
	err := ProtoBuf.Bind(req, &obj)
	assert.Error(t, err)
}

func TestValidationDisabled(t *testing.T) {
	backup := Validator
	Validator = nil
	defer func() { Validator = backup }()

	var obj FooStruct
	req := requestWithBody("POST", "/", `{"bar": "foo"}`)
	err := ProtoBuf.Bind(req, &obj)
	assert.NoError(t, err)
}

type HogeStruct struct {
	Hoge int32 `json:"hoge" binding:"exists" protobuf:"varint,1,opt,json=hoge,proto3"`
}

func (fd *HogeStruct) Reset()         {}
func (fd *HogeStruct) String() string { return "" }
func (fd *HogeStruct) ProtoMessage()  {}

func TestExistsSucceeds(t *testing.T) {
	var obj HogeStruct
	req := requestWithBody("POST", "/", `{"hoge": 0}`)
	err := ProtoBuf.Bind(req, &obj)
	assert.NoError(t, err)
}

func TestExistsFails(t *testing.T) {
	type HogeStruct2 struct {
		Hoge int `json:"foo" binding:"exists" protobuf:"varint,1,opt,json=hoge,proto3"`
	}

	var obj HogeStruct2
	req := requestWithBody("POST", "/", `{"boen": 0}`)
	err := ProtoBuf.Bind(req, &obj)
	assert.Error(t, err)
}

func testFormBindingEmbeddedStruct(t *testing.T, method, path, badPath, body, badBody string) {
	b := Form
	assert.Equal(t, "form", b.Name())

	obj := QueryTest{}
	req := requestWithBody(method, path, body)
	if method == "POST" {
		req.Header.Add("Content-Type", MIMEPOSTForm)
	}
	err := b.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, 1, obj.Page)
	assert.Equal(t, 2, obj.Size)
	assert.Equal(t, "test-appkey", obj.Appkey)

}

func testFormBinding(t *testing.T, method, path, badPath, body, badBody string) {
	b := Form
	assert.Equal(t, "form", b.Name())

	obj := FooBarStruct{}
	req := requestWithBody(method, path, body)
	if method == "POST" {
		req.Header.Add("Content-Type", MIMEPOSTForm)
	}
	err := b.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, "foo", obj.Bar)

	obj = FooBarStruct{}
	req = requestWithBody(method, badPath, badBody)
	err = ProtoBuf.Bind(req, &obj)
	assert.Error(t, err)
}

func testFormBindingDefaultValue(t *testing.T, method, path, badPath, body, badBody string) {
	b := Form
	assert.Equal(t, "form", b.Name())

	obj := FooDefaultBarStruct{}
	req := requestWithBody(method, path, body)
	if method == "POST" {
		req.Header.Add("Content-Type", MIMEPOSTForm)
	}
	err := b.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, "hello", obj.Bar)

	obj = FooDefaultBarStruct{}
	req = requestWithBody(method, badPath, badBody)
	err = ProtoBuf.Bind(req, &obj)
	assert.Error(t, err)
}

func TestFormBindingFail(t *testing.T) {
	b := Form
	assert.Equal(t, "form", b.Name())

	obj := FooBarStruct{}
	req, _ := http.NewRequest("POST", "/", nil)
	err := b.Bind(req, &obj)
	assert.Error(t, err)
}

func TestFormBindingMultipartFail(t *testing.T) {
	obj := FooBarStruct{}
	req, err := http.NewRequest("POST", "/", strings.NewReader("foo=bar"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+";boundary=testboundary")
	_, err = req.MultipartReader()
	assert.NoError(t, err)
	err = Form.Bind(req, &obj)
	assert.Error(t, err)
}

func TestFormPostBindingFail(t *testing.T) {
	b := FormPost
	assert.Equal(t, "form-urlencoded", b.Name())

	obj := FooBarStruct{}
	req, _ := http.NewRequest("POST", "/", nil)
	err := b.Bind(req, &obj)
	assert.Error(t, err)
}

func TestFormMultipartBindingFail(t *testing.T) {
	b := FormMultipart
	assert.Equal(t, "multipart/form-data", b.Name())

	obj := FooBarStruct{}
	req, _ := http.NewRequest("POST", "/", nil)
	err := b.Bind(req, &obj)
	assert.Error(t, err)
}
func testFormBindingIgnoreField(t *testing.T, method, path, badPath, body, badBody string) {
	b := Form
	assert.Equal(t, "form", b.Name())

	obj := FooStructForIgnoreFormTag{}
	req := requestWithBody(method, path, body)
	if method == "POST" {
		req.Header.Add("Content-Type", MIMEPOSTForm)
	}
	err := b.Bind(req, &obj)
	assert.NoError(t, err)

	assert.Nil(t, obj.Foo)
}

func testFormBindingForType(t *testing.T, method, path, badPath, body, badBody string, typ string) {
	b := Form
	assert.Equal(t, "form", b.Name())

	req := requestWithBody(method, path, body)
	if method == "POST" {
		req.Header.Add("Content-Type", MIMEPOSTForm)
	}
	switch typ {
	case "Slice":
		obj := FooStructForSliceType{}
		err := b.Bind(req, &obj)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2}, obj.SliceFoo)

		obj = FooStructForSliceType{}
		req = requestWithBody(method, badPath, badBody)
		err = ProtoBuf.Bind(req, &obj)
		assert.Error(t, err)
	case "Struct":
		obj := FooStructForStructType{}
		err := b.Bind(req, &obj)
		assert.NoError(t, err)
		assert.Equal(t,
			123,
			obj.StructFoo.Idx)
	case "StructPointer":
		obj := FooStructForStructPointerType{}
		err := b.Bind(req, &obj)
		assert.NoError(t, err)
		assert.Equal(t,
			"thinkerou",
			obj.StructPointerFoo.Name)
	case "Map":
		obj := FooStructForMapType{}
		err := b.Bind(req, &obj)
		assert.NoError(t, err)
		assert.Equal(t, float64(123), obj.MapFoo["bar"].(float64))
	case "SliceMap":
		obj := FooStructForSliceMapType{}
		err := b.Bind(req, &obj)
		assert.Error(t, err)
	case "Ptr":
		obj := FooStructForStringPtrType{}
		err := b.Bind(req, &obj)
		assert.NoError(t, err)
		assert.Nil(t, obj.PtrFoo)
		assert.Equal(t, "test", *obj.PtrBar)

		obj = FooStructForStringPtrType{}
		obj.PtrBar = new(string)
		err = b.Bind(req, &obj)
		assert.NoError(t, err)
		assert.Equal(t, "test", *obj.PtrBar)

		objErr := FooStructForMapPtrType{}
		err = b.Bind(req, &objErr)
		assert.Error(t, err)

		obj = FooStructForStringPtrType{}
		req = requestWithBody(method, badPath, badBody)
		err = b.Bind(req, &obj)
		assert.Error(t, err)
	}
}

func testQueryBinding(t *testing.T, method, path, badPath, body, badBody string) {
	b := Query
	assert.Equal(t, "query", b.Name())

	obj := FooBarStruct{}
	req := requestWithBody(method, path, body)
	if method == "POST" {
		req.Header.Add("Content-Type", MIMEPOSTForm)
	}
	err := b.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, "foo", obj.Bar)
}

func testQueryBindingFail(t *testing.T, method, path, badPath, body, badBody string) {
	b := Query
	assert.Equal(t, "query", b.Name())

	obj := FooStructForMapType{}
	req := requestWithBody(method, path, body)
	if method == "POST" {
		req.Header.Add("Content-Type", MIMEPOSTForm)
	}
	err := b.Bind(req, &obj)
	assert.Error(t, err)
}

func testQueryBindingBoolFail(t *testing.T, method, path, badPath, body, badBody string) {
	b := Query
	assert.Equal(t, "query", b.Name())

	obj := FooStructForBoolType{}
	req := requestWithBody(method, path, body)
	if method == "POST" {
		req.Header.Add("Content-Type", MIMEPOSTForm)
	}
	err := b.Bind(req, &obj)
	assert.Error(t, err)
}

func testBodyBinding(t *testing.T, b Binding, name, path, badPath, body, badBody string) {
	assert.Equal(t, name, b.Name())

	obj := FooStruct{}
	req := requestWithBody("POST", path, body)
	err := b.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "bar", obj.Foo)

	obj = FooStruct{}
	req = requestWithBody("POST", badPath, badBody)
	err = ProtoBuf.Bind(req, &obj)
	assert.Error(t, err)
}

type hook struct{}

func (h hook) Read([]byte) (int, error) {
	return 0, errors.New("error")
}

func testProtoBodyBindingFail(t *testing.T, b Binding, name, path, badPath, body, badBody string) {
	assert.Equal(t, name, b.Name())

	obj := protoexample.Test{}
	req := requestWithBody("POST", path, body)

	req.Body = ioutil.NopCloser(&hook{})
	req.Header.Add("Content-Type", MIMEPROTOBUF)
	err := b.Bind(req, &obj)
	assert.Error(t, err)

	obj = protoexample.Test{}
	req = requestWithBody("POST", badPath, badBody)
	req.Header.Add("Content-Type", MIMEPROTOBUF)
	err = ProtoBuf.Bind(req, &obj)
	assert.Error(t, err)
}

func requestWithBody(method, path, body string) (req *http.Request) {
	req, _ = http.NewRequest(method, path, bytes.NewBufferString(body))
	return
}
