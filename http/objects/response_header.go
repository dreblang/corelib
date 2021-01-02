package objects

import (
	"fmt"

	"github.com/dreblang/core/object"
	"github.com/valyala/fasthttp"
)

const ResponseHeaderObj = "http:ResponseHeader"

type ResponseHeader struct {
	header *fasthttp.ResponseHeader
}

func NewResponseHeader(res *fasthttp.Response) *ResponseHeader {
	return &ResponseHeader{
		header: &res.Header,
	}
}

func (obj *ResponseHeader) Type() object.ObjectType { return ResponseHeaderObj }
func (obj *ResponseHeader) Inspect() string {
	return fmt.Sprintf("http:ResponseHeader[%p]", obj)
}

func (obj *ResponseHeader) String() string { return "ResponseHeader" }

func (obj *ResponseHeader) GetMember(name string) object.Object {
	switch name {
	case "add":
		return &object.MemberFn{
			Obj: obj,
			Fn:  responseHeaderAdd,
		}

	case "set":
		return &object.MemberFn{
			Obj: obj,
			Fn:  responseHeaderSet,
		}

	case "del":
		return &object.MemberFn{
			Obj: obj,
			Fn:  responseHeaderDel,
		}
	}
	return object.NewError("No member named [%s]", name)
}

func (obj *ResponseHeader) SetMember(name string, value object.Object) object.Object {
	return object.NewError("No member named [%s]", name)
}

func (obj *ResponseHeader) Equals(other object.Object) bool {
	return false
}

func responseHeaderAdd(this object.Object, args ...object.Object) object.Object {
	rh := this.(*ResponseHeader)
	if len(args) == 2 {
		if key, ok := args[0].(*object.String); ok {
			if value, ok := args[1].(*object.String); ok {
				rh.header.Add(key.Value, value.Value)
				return object.NullObject
			}
		}
	}
	return object.NewError("Could not execute header add operation. Invalid arguments!")
}

func responseHeaderSet(this object.Object, args ...object.Object) object.Object {
	rh := this.(*ResponseHeader)
	if len(args) == 2 {
		if key, ok := args[0].(*object.String); ok {
			if value, ok := args[1].(*object.String); ok {
				rh.header.Set(key.Value, value.Value)
				return object.NullObject
			}
		}
	}
	return object.NewError("Could not execute header set operation. Invalid arguments!")
}

func responseHeaderDel(this object.Object, args ...object.Object) object.Object {
	rh := this.(*ResponseHeader)
	if len(args) == 1 {
		if key, ok := args[0].(*object.String); ok {
			rh.header.Del(key.Value)
			return object.NullObject
		}
	}
	return object.NewError("Could not execute header set operation. Invalid arguments!")
}
