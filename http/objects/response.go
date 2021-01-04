package objects

import (
	"fmt"

	"github.com/dreblang/core/object"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

const ResponseObj = "http:Response"

type Response struct {
	ctx     *fasthttp.Response
	headers *ResponseHeader
	_dict   map[string]object.Object
}

func NewResponse(ctx *routing.Context) *Response {
	resp := &Response{
		ctx:     &ctx.Response,
		headers: NewResponseHeader(&ctx.Response),
	}

	resp.makeDict()

	return resp
}

func (obj *Response) makeDict() {
	obj._dict = map[string]object.Object{}
	obj._dict["headers"] = obj.headers
	obj._dict["status"] = &object.MemberFn{
		Obj: obj,
		Fn:  responseStatus,
	}
}

func (obj *Response) Type() object.ObjectType { return ResponseObj }
func (obj *Response) Inspect() string {
	return fmt.Sprintf("http:Response[%p]", obj)
}

func (obj *Response) String() string { return "Response" }

func (obj *Response) GetMember(name string) object.Object {
	if mem, ok := obj._dict[name]; ok {
		return mem
	}
	return object.NewError("No member named [%s]", name)
}

func (obj *Response) SetMember(name string, value object.Object) object.Object {
	return object.NewError("No member named [%s]", name)
}

func (obj *Response) Equals(other object.Object) bool {
	return false
}

func responseStatus(this object.Object, args ...object.Object) object.Object {
	r := this.(*Response)
	if len(args) == 1 {
		if value, ok := args[0].(*object.Integer); ok {
			r.ctx.SetStatusCode(int(value.Value))
		}
	}
	return object.NewError("Could not set status. Invalid arguments!")
}

func responseRedirect(this object.Object, args ...object.Object) object.Object {
	r := this.(*Response)
	if len(args) == 1 {
		if value, ok := args[0].(*object.String); ok {
			r.ctx.SetStatusCode(fasthttp.StatusFound)
			r.ctx.Header.Set("location", value.Value)
		}
	}
	return object.NewError("Could not set status. Invalid arguments!")
}
