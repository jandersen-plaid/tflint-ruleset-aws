// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule checks the pattern is valid
type AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSagemakerUserProfileInvalidSingleSignOnUserValueRule returns new rule with default attributes
func NewAwsSagemakerUserProfileInvalidSingleSignOnUserValueRule() *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule {
	return &AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule{
		resourceType:  "aws_sagemaker_user_profile",
		attributeName: "single_sign_on_user_value",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Name() string {
	return "aws_sagemaker_user_profile_invalid_single_sign_on_user_value"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"single_sign_on_user_value must be 256 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
