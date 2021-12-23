// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyAppInvalidPlatformRule checks the pattern is valid
type AwsAmplifyAppInvalidPlatformRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAmplifyAppInvalidPlatformRule returns new rule with default attributes
func NewAwsAmplifyAppInvalidPlatformRule() *AwsAmplifyAppInvalidPlatformRule {
	return &AwsAmplifyAppInvalidPlatformRule{
		resourceType:  "aws_amplify_app",
		attributeName: "platform",
		enum: []string{
			"WEB",
		},
	}
}

// Name returns the rule name
func (r *AwsAmplifyAppInvalidPlatformRule) Name() string {
	return "aws_amplify_app_invalid_platform"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyAppInvalidPlatformRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyAppInvalidPlatformRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyAppInvalidPlatformRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyAppInvalidPlatformRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as platform`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
