// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationSmbInvalidSubdirectoryRule checks the pattern is valid
type AwsDatasyncLocationSmbInvalidSubdirectoryRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationSmbInvalidSubdirectoryRule returns new rule with default attributes
func NewAwsDatasyncLocationSmbInvalidSubdirectoryRule() *AwsDatasyncLocationSmbInvalidSubdirectoryRule {
	return &AwsDatasyncLocationSmbInvalidSubdirectoryRule{
		resourceType:  "aws_datasync_location_smb",
		attributeName: "subdirectory",
		max:           4096,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Name() string {
	return "aws_datasync_location_smb_invalid_subdirectory"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"subdirectory must be 4096 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
