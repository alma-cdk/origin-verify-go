package almacdkoriginverify

import (
	_init_ "github.com/alma-cdk/origin-verify-go/almacdkoriginverify/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alma-cdk/origin-verify-go/almacdkoriginverify/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates an origin with WAFv2 WebACL to verify traffic contains specific header with a secret value.
// Experimental.
type OriginVerify interface {
	constructs.Construct
	IVerification
	// CloudFront Origin Custom Header name used in the WAFv2 WebACL verification.
	// Default: 'x-origin-verify'.
	//
	// Experimental.
	HeaderName() *string
	// Secret Value used as the CloudFront Origin Custom Header value.
	//
	// Example:
	//   'xxxxEXAMPLESECRET'
	//
	// Experimental.
	HeaderValue() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OriginVerify
type jsiiProxy_OriginVerify struct {
	internal.Type__constructsConstruct
	jsiiProxy_IVerification
}

func (j *jsiiProxy_OriginVerify) HeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginVerify) HeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginVerify) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Associates an origin with WAFv2 WebACL to verify traffic contains specific header with a secret value.
//
// Use `verifyHeader` value to assign custom headers into CloudFront config.
//
// Example:
//   import { OriginVerify } from '@alma-cdk/origin-verify';
//   import { Distribution } from 'aws-cdk-lib/aws-cloudfront';
//
//   const api: RestApi; // TODO: implement the RestApi
//   const apiDomain: string; // TODO: implement the domain
//
//   const verification = new OriginVerify(this, 'OriginVerify', {
//     origin: api.deploymentStage,
//   });
//
//   new Distribution(this, 'CDN', {
//     defaultBehavior: {
//       origin: new HttpOrigin(apiDomain, {
//         customHeaders: {
//           [verification.headerName]: verification.headerValue,
//         },
//         protocolPolicy: OriginProtocolPolicy.HTTPS_ONLY,
//       })
//     },
//   })
//
// Experimental.
func NewOriginVerify(scope constructs.Construct, id *string, props *OriginVerifyProps) OriginVerify {
	_init_.Initialize()

	if err := validateNewOriginVerifyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OriginVerify{}

	_jsii_.Create(
		"@alma-cdk/origin-verify.OriginVerify",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Associates an origin with WAFv2 WebACL to verify traffic contains specific header with a secret value.
//
// Use `verifyHeader` value to assign custom headers into CloudFront config.
//
// Example:
//   import { OriginVerify } from '@alma-cdk/origin-verify';
//   import { Distribution } from 'aws-cdk-lib/aws-cloudfront';
//
//   const api: RestApi; // TODO: implement the RestApi
//   const apiDomain: string; // TODO: implement the domain
//
//   const verification = new OriginVerify(this, 'OriginVerify', {
//     origin: api.deploymentStage,
//   });
//
//   new Distribution(this, 'CDN', {
//     defaultBehavior: {
//       origin: new HttpOrigin(apiDomain, {
//         customHeaders: {
//           [verification.headerName]: verification.headerValue,
//         },
//         protocolPolicy: OriginProtocolPolicy.HTTPS_ONLY,
//       })
//     },
//   })
//
// Experimental.
func NewOriginVerify_Override(o OriginVerify, scope constructs.Construct, id *string, props *OriginVerifyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/origin-verify.OriginVerify",
		[]interface{}{scope, id, props},
		o,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func OriginVerify_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOriginVerify_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/origin-verify.OriginVerify",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OriginVerify_OriginVerifyHeader() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@alma-cdk/origin-verify.OriginVerify",
		"OriginVerifyHeader",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OriginVerify) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

