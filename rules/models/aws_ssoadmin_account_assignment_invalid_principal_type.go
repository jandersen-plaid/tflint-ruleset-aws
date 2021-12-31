// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule checks the pattern is valid
type AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsoadminAccountAssignmentInvalidPrincipalTypeRule returns new rule with default attributes
func NewAwsSsoadminAccountAssignmentInvalidPrincipalTypeRule() *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule {
	return &AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule{
		resourceType:  "aws_ssoadmin_account_assignment",
		attributeName: "principal_type",
		enum: []string{
			"USER",
			"GROUP",
		},
	}
}

// Name returns the rule name
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Name() string {
	return "aws_ssoadmin_account_assignment_invalid_principal_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as principal_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}