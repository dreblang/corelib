package objects

import (
	"fmt"

	"github.com/dreblang/core/object"
	routing "github.com/qiangxue/fasthttp-routing"
)

const ParamsObj = "http:Params"

type Params struct {
	ctx *routing.Context
}

func NewParams(ctx *routing.Context) *Params {
	return &Params{ctx: ctx}
}

func (obj *Params) Type() object.ObjectType { return ParamsObj }
func (obj *Params) Inspect() string {
	return fmt.Sprintf("http:Params[%p]", obj)
}

func (obj *Params) String() string { return "Params" }

func (obj *Params) GetMember(name string) object.Object {
	return &object.String{Value: obj.ctx.Param(name)}
}

func (obj *Params) SetMember(name string, value object.Object) object.Object {
	return object.NewError("No member named [%s]", name)
}

func (obj *Params) Equals(other object.Object) bool {
	return false
}
