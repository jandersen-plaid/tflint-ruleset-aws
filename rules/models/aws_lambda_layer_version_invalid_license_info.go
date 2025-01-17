// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaLayerVersionInvalidLicenseInfoRule checks the pattern is valid
type AwsLambdaLayerVersionInvalidLicenseInfoRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsLambdaLayerVersionInvalidLicenseInfoRule returns new rule with default attributes
func NewAwsLambdaLayerVersionInvalidLicenseInfoRule() *AwsLambdaLayerVersionInvalidLicenseInfoRule {
	return &AwsLambdaLayerVersionInvalidLicenseInfoRule{
		resourceType:  "aws_lambda_layer_version",
		attributeName: "license_info",
		max:           512,
	}
}

// Name returns the rule name
func (r *AwsLambdaLayerVersionInvalidLicenseInfoRule) Name() string {
	return "aws_lambda_layer_version_invalid_license_info"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaLayerVersionInvalidLicenseInfoRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaLayerVersionInvalidLicenseInfoRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaLayerVersionInvalidLicenseInfoRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaLayerVersionInvalidLicenseInfoRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"license_info must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
