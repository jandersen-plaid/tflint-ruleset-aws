// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightUserInvalidNamespaceRule checks the pattern is valid
type AwsQuicksightUserInvalidNamespaceRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsQuicksightUserInvalidNamespaceRule returns new rule with default attributes
func NewAwsQuicksightUserInvalidNamespaceRule() *AwsQuicksightUserInvalidNamespaceRule {
	return &AwsQuicksightUserInvalidNamespaceRule{
		resourceType:  "aws_quicksight_user",
		attributeName: "namespace",
		max:           64,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9._-]*$`),
	}
}

// Name returns the rule name
func (r *AwsQuicksightUserInvalidNamespaceRule) Name() string {
	return "aws_quicksight_user_invalid_namespace"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightUserInvalidNamespaceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightUserInvalidNamespaceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightUserInvalidNamespaceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightUserInvalidNamespaceRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"namespace must be 64 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9._-]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}