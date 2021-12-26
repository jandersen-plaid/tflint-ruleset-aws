// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerDomainInvalidAuthModeRule checks the pattern is valid
type AwsSagemakerDomainInvalidAuthModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSagemakerDomainInvalidAuthModeRule returns new rule with default attributes
func NewAwsSagemakerDomainInvalidAuthModeRule() *AwsSagemakerDomainInvalidAuthModeRule {
	return &AwsSagemakerDomainInvalidAuthModeRule{
		resourceType:  "aws_sagemaker_domain",
		attributeName: "auth_mode",
		enum: []string{
			"SSO",
			"IAM",
		},
	}
}

// Name returns the rule name
func (r *AwsSagemakerDomainInvalidAuthModeRule) Name() string {
	return "aws_sagemaker_domain_invalid_auth_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerDomainInvalidAuthModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerDomainInvalidAuthModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerDomainInvalidAuthModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerDomainInvalidAuthModeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as auth_mode`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
