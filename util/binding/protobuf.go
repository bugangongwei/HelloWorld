// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	errNotProtoMessageType = errors.New("obj not proto message type ")
	errNilProtoMessage     = errors.New("obj is nil")
	errNilRequest          = errors.New("reqeust or request body is nil")
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return errNilRequest
	}
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return b.BindBody(buf, obj)
}

// BindBody .
func (protobufBinding) BindBody(body []byte, obj interface{}) error {
	if obj == nil {
		return errNilProtoMessage
	}
	pm, ok := obj.(proto.Message)
	if !ok {
		return errNotProtoMessageType
	}
	m := jsonpb.Unmarshaler{AllowUnknownFields: true}
	err := m.Unmarshal(bytes.NewReader(body), pm)
	if err != nil {
		return err
	}
	return validate(obj)
}
