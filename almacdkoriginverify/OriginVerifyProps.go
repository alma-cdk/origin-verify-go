package almacdkoriginverify

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for `OriginVerify` constructor.
// Experimental.
type OriginVerifyProps struct {
	// Origin to protect.
	//
	// Accepted types:
	// - `IStage` (from `aws-cdk-lib/aws-apigateway`)
	// - `IApplicationLoadBalancer` (from `aws-cdk-lib/aws-elasticloadbalancingv2`).
	// Experimental.
	Origin interface{} `field:"required" json:"origin" yaml:"origin"`
	// Metric name for the WebACL.
	// Default: 'OriginVerifyWebAcl'.
	//
	// Experimental.
	AclMetricName *string `field:"optional" json:"aclMetricName" yaml:"aclMetricName"`
	// By default `x-origin-verify` is used.
	//
	// To override it, provide a value for
	// this. Recommendation is to use something with a `x-` prefix.
	// Default: 'x-origin-verify'.
	//
	// Experimental.
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// Metric name for the allowed requests.
	// Default: 'OriginVerifyAllowedRequests'.
	//
	// Experimental.
	RuleMetricName *string `field:"optional" json:"ruleMetricName" yaml:"ruleMetricName"`
	// Any additional rules to add into the created WAFv2 WebACL.
	// Experimental.
	Rules *[]interface{} `field:"optional" json:"rules" yaml:"rules"`
	// The secret which is used to verify the CloudFront distribution.
	//
	// Optional: By default this construct will generate a `new Secret`.
	// Default: new Secret().secretValue
	//
	// Experimental.
	SecretValue awscdk.SecretValue `field:"optional" json:"secretValue" yaml:"secretValue"`
}

