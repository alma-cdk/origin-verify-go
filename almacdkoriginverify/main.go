// Enforce origin traffic via CloudFront.
package almacdkoriginverify

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"@alma-cdk/origin-verify.IVerification",
		reflect.TypeOf((*IVerification)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
		},
		func() interface{} {
			return &jsiiProxy_IVerification{}
		},
	)
	_jsii_.RegisterClass(
		"@alma-cdk/origin-verify.OriginVerify",
		reflect.TypeOf((*OriginVerify)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerValue", GoGetter: "HeaderValue"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OriginVerify{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVerification)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/origin-verify.OriginVerifyProps",
		reflect.TypeOf((*OriginVerifyProps)(nil)).Elem(),
	)
}
