// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyAppInvalidDescriptionRule checks the pattern is valid
type AwsAmplifyAppInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsAmplifyAppInvalidDescriptionRule returns new rule with default attributes
func NewAwsAmplifyAppInvalidDescriptionRule() *AwsAmplifyAppInvalidDescriptionRule {
	return &AwsAmplifyAppInvalidDescriptionRule{
		resourceType:  "aws_amplify_app",
		attributeName: "description",
		max:           1000,
	}
}

// Name returns the rule name
func (r *AwsAmplifyAppInvalidDescriptionRule) Name() string {
	return "aws_amplify_app_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyAppInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyAppInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyAppInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyAppInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 1000 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
