// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamFleetInvalidDisplayNameRule checks the pattern is valid
type AwsAppstreamFleetInvalidDisplayNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsAppstreamFleetInvalidDisplayNameRule returns new rule with default attributes
func NewAwsAppstreamFleetInvalidDisplayNameRule() *AwsAppstreamFleetInvalidDisplayNameRule {
	return &AwsAppstreamFleetInvalidDisplayNameRule{
		resourceType:  "aws_appstream_fleet",
		attributeName: "display_name",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Name() string {
	return "aws_appstream_fleet_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamFleetInvalidDisplayNameRule) Check(runner tflint.Runner) error {
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
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"display_name must be 100 characters or less",
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
