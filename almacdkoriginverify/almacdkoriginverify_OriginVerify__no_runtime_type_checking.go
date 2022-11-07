//go:build no_runtime_type_checking

// Enforce origin traffic via CloudFront.
package almacdkoriginverify

// Building without runtime type checking enabled, so all the below just return nil

func validateOriginVerify_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewOriginVerifyParameters(scope constructs.Construct, id *string, props *OriginVerifyProps) error {
	return nil
}

