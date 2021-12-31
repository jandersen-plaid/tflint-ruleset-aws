// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule checks the pattern is valid
type AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule returns new rule with default attributes
func NewAwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule() *AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule {
	return &AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule{
		resourceType:  "aws_datasync_location_fsx_windows_file_system",
		attributeName: "password",
		max:           104,
		pattern:       regexp.MustCompile(`^.{0,104}$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule) Name() string {
	return "aws_datasync_location_fsx_windows_file_system_invalid_password"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidPasswordRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"password must be 104 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					`password does not match valid pattern ^.{0,104}$`,
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
