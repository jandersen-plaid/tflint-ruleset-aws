// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsKmsKeyInvalidKeyUsageRule checks the pattern is valid
type AwsKmsKeyInvalidKeyUsageRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsKmsKeyInvalidKeyUsageRule returns new rule with default attributes
func NewAwsKmsKeyInvalidKeyUsageRule() *AwsKmsKeyInvalidKeyUsageRule {
	return &AwsKmsKeyInvalidKeyUsageRule{
		resourceType:  "aws_kms_key",
		attributeName: "key_usage",
		enum: []string{
			"SIGN_VERIFY",
			"ENCRYPT_DECRYPT",
		},
	}
}

// Name returns the rule name
func (r *AwsKmsKeyInvalidKeyUsageRule) Name() string {
	return "aws_kms_key_invalid_key_usage"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsKeyInvalidKeyUsageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKmsKeyInvalidKeyUsageRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsKeyInvalidKeyUsageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsKeyInvalidKeyUsageRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as key_usage`, truncateLongMessage(val)),
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
