// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcsAccountSettingDefaultInvalidNameRule checks the pattern is valid
type AwsEcsAccountSettingDefaultInvalidNameRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEcsAccountSettingDefaultInvalidNameRule returns new rule with default attributes
func NewAwsEcsAccountSettingDefaultInvalidNameRule() *AwsEcsAccountSettingDefaultInvalidNameRule {
	return &AwsEcsAccountSettingDefaultInvalidNameRule{
		resourceType:  "aws_ecs_account_setting_default",
		attributeName: "name",
		enum: []string{
			"serviceLongArnFormat",
			"taskLongArnFormat",
			"containerInstanceLongArnFormat",
			"awsvpcTrunking",
			"containerInsights",
		},
	}
}

// Name returns the rule name
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Name() string {
	return "aws_ecs_account_setting_default_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcsAccountSettingDefaultInvalidNameRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as name`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}