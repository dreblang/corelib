package http

import (
	"github.com/dreblang/core/object"
	"github.com/dreblang/corelib/http/objects"
)

var httpScope = &object.Scope{
	Name: "http",
	Exports: map[string]object.Object{
		"StatusContinue":           &object.Integer{Value: 100},
		"StatusSwitchingProtocols": &object.Integer{Value: 101},
		"StatusProcessing":         &object.Integer{Value: 102},
		"StatusEarlyHints":         &object.Integer{Value: 103},

		"StatusOK":                   &object.Integer{Value: 200},
		"StatusCreated":              &object.Integer{Value: 201},
		"StatusAccepted":             &object.Integer{Value: 202},
		"StatusNonAuthoritativeInfo": &object.Integer{Value: 203},
		"StatusNoContent":            &object.Integer{Value: 204},
		"StatusResetContent":         &object.Integer{Value: 205},
		"StatusPartialContent":       &object.Integer{Value: 206},
		"StatusMultiStatus":          &object.Integer{Value: 207},
		"StatusAlreadyReported":      &object.Integer{Value: 208},
		"StatusIMUsed":               &object.Integer{Value: 226},

		"StatusMultipleChoices":   &object.Integer{Value: 300},
		"StatusMovedPermanently":  &object.Integer{Value: 301},
		"StatusFound":             &object.Integer{Value: 302},
		"StatusSeeOther":          &object.Integer{Value: 303},
		"StatusNotModified":       &object.Integer{Value: 304},
		"StatusUseProxy":          &object.Integer{Value: 305},
		"StatusTemporaryRedirect": &object.Integer{Value: 307},
		"StatusPermanentRedirect": &object.Integer{Value: 308},

		"StatusBadRequest":                   &object.Integer{Value: 400},
		"StatusUnauthorized":                 &object.Integer{Value: 401},
		"StatusPaymentRequired":              &object.Integer{Value: 402},
		"StatusForbidden":                    &object.Integer{Value: 403},
		"StatusNotFound":                     &object.Integer{Value: 404},
		"StatusMethodNotAllowed":             &object.Integer{Value: 405},
		"StatusNotAcceptable":                &object.Integer{Value: 406},
		"StatusProxyAuthRequired":            &object.Integer{Value: 407},
		"StatusRequestTimeout":               &object.Integer{Value: 408},
		"StatusConflict":                     &object.Integer{Value: 409},
		"StatusGone":                         &object.Integer{Value: 410},
		"StatusLengthRequired":               &object.Integer{Value: 411},
		"StatusPreconditionFailed":           &object.Integer{Value: 412},
		"StatusRequestEntityTooLarge":        &object.Integer{Value: 413},
		"StatusRequestURITooLong":            &object.Integer{Value: 414},
		"StatusUnsupportedMediaType":         &object.Integer{Value: 415},
		"StatusRequestedRangeNotSatisfiable": &object.Integer{Value: 416},
		"StatusExpectationFailed":            &object.Integer{Value: 417},
		"StatusTeapot":                       &object.Integer{Value: 418},
		"StatusMisdirectedRequest":           &object.Integer{Value: 421},
		"StatusUnprocessableEntity":          &object.Integer{Value: 422},
		"StatusLocked":                       &object.Integer{Value: 423},
		"StatusFailedDependency":             &object.Integer{Value: 424},
		"StatusUpgradeRequired":              &object.Integer{Value: 426},
		"StatusPreconditionRequired":         &object.Integer{Value: 428},
		"StatusTooManyRequests":              &object.Integer{Value: 429},
		"StatusRequestHeaderFieldsTooLarge":  &object.Integer{Value: 431},
		"StatusUnavailableForLegalReasons":   &object.Integer{Value: 451},

		"StatusInternalServerError":           &object.Integer{Value: 500},
		"StatusNotImplemented":                &object.Integer{Value: 501},
		"StatusBadGateway":                    &object.Integer{Value: 502},
		"StatusServiceUnavailable":            &object.Integer{Value: 503},
		"StatusGatewayTimeout":                &object.Integer{Value: 504},
		"StatusHTTPVersionNotSupported":       &object.Integer{Value: 505},
		"StatusVariantAlsoNegotiates":         &object.Integer{Value: 506},
		"StatusInsufficientStorage":           &object.Integer{Value: 507},
		"StatusLoopDetected":                  &object.Integer{Value: 508},
		"StatusNotExtended":                   &object.Integer{Value: 510},
		"StatusNetworkAuthenticationRequired": &object.Integer{Value: 511},

		"createServer": &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				return objects.NewServer()
			},
		},
	},
}

// Load the scope
func Load() *object.Scope {
	return httpScope
}
