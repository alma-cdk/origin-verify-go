// Enforce origin traffic via CloudFront.
package almacdkoriginverify

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface describing the "contract" of return values from the constructor.
// Experimental.
type IVerification interface {
	// CloudFront Origin Custom Header name used in the WAFv2 WebACL verification.
	// Experimental.
	HeaderName() *string
	// Secret Value used as the CloudFront Origin Custom Header value.
	//
	// Example:
	//   'xxxxEXAMPLESECRET'
	//
	// Experimental.
	HeaderValue() *string
}

// The jsii proxy for IVerification
type jsiiProxy_IVerification struct {
	_ byte // padding
}

func (j *jsiiProxy_IVerification) HeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerification) HeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerValue",
		&returns,
	)
	return returns
}

