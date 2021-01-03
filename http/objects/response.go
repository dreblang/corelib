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
}

func NewResponse(ctx *routing.Context) *Response {
	return &Response{
		ctx:     &ctx.Response,
		headers: NewResponseHeader(&ctx.Response),
	}
}

func (obj *Response) Type() object.ObjectType { return ResponseObj }
func (obj *Response) Inspect() string {
	return fmt.Sprintf("http:Response[%p]", obj)
}

func (obj *Response) String() string { return "Response" }

func (obj *Response) GetMember(name string) object.Object {
	switch name {
	case "headers":
		return obj.headers

	case "status":
		return &object.MemberFn{
			Obj: obj,
			Fn:  responseStatus,
		}
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
