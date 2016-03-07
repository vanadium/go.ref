// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: channel.vdl

package channel

import (
	// VDL system imports
	"v.io/v23/vdl"
	"v.io/v23/vom"
)

type Request struct {
	Type string
	Seq  uint32
	Body *vom.RawBytes
}

func (Request) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/channel.Request"`
}) {
}

func (m *Request) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Request == nil || __VDLTypechannel0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Type")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Type), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Seq")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromUint(uint64(m.Seq), vdl.Uint32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Body")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if m.Body == nil {
			if err := fieldTarget7.FromNil(vdl.AnyType); err != nil {
				return err
			}
		} else {
			if err := m.Body.FillVDLTarget(fieldTarget7, vdl.AnyType); err != nil {
				return err
			}
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Request) MakeVDLTarget() vdl.Target {
	return nil
}

type Response struct {
	ReqSeq uint32
	Err    string // TODO(bprosnitz) change this back to error when it is possible to do so. (issue 368)
	Body   *vom.RawBytes
}

func (Response) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/wspr/internal/channel.Response"`
}) {
}

func (m *Response) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Response == nil || __VDLTypechannel1 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("ReqSeq")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromUint(uint64(m.ReqSeq), vdl.Uint32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Err")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.Err), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Body")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if m.Body == nil {
			if err := fieldTarget7.FromNil(vdl.AnyType); err != nil {
				return err
			}
		} else {
			if err := m.Body.FillVDLTarget(fieldTarget7, vdl.AnyType); err != nil {
				return err
			}
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Response) MakeVDLTarget() vdl.Target {
	return nil
}

type (
	// Message represents any single field of the Message union type.
	Message interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the Message union type.
		__VDLReflect(__MessageReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
	}
	// MessageRequest represents field Request of the Message union type.
	MessageRequest struct{ Value Request }
	// MessageResponse represents field Response of the Message union type.
	MessageResponse struct{ Value Response }
	// __MessageReflect describes the Message union type.
	__MessageReflect struct {
		Name  string `vdl:"v.io/x/ref/services/wspr/internal/channel.Message"`
		Type  Message
		Union struct {
			Request  MessageRequest
			Response MessageResponse
		}
	}
)

func (x MessageRequest) Index() int                    { return 0 }
func (x MessageRequest) Interface() interface{}        { return x.Value }
func (x MessageRequest) Name() string                  { return "Request" }
func (x MessageRequest) __VDLReflect(__MessageReflect) {}

func (m MessageRequest) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Message)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Request")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Request); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m MessageRequest) MakeVDLTarget() vdl.Target {
	return nil
}

func (x MessageResponse) Index() int                    { return 1 }
func (x MessageResponse) Interface() interface{}        { return x.Value }
func (x MessageResponse) Name() string                  { return "Response" }
func (x MessageResponse) __VDLReflect(__MessageReflect) {}

func (m MessageResponse) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Message)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Response")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Response); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m MessageResponse) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))
	vdl.Register((*Message)(nil))
}

var __VDLTypechannel0 *vdl.Type = vdl.TypeOf((*Request)(nil))
var __VDLTypechannel1 *vdl.Type = vdl.TypeOf((*Response)(nil))
var __VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Message *vdl.Type = vdl.TypeOf(Message(MessageRequest{Request{}}))
var __VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Request *vdl.Type = vdl.TypeOf(Request{})
var __VDLType_channel_v_io_x_ref_services_wspr_internal_channel_Response *vdl.Type = vdl.TypeOf(Response{})

func __VDLEnsureNativeBuilt_channel() {
}
