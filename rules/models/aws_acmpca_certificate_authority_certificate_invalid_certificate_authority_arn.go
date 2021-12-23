// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule checks the pattern is valid
type AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule returns new rule with default attributes
func NewAwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule() *AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule {
	return &AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule{
		resourceType:  "aws_acmpca_certificate_authority_certificate",
		attributeName: "certificate_authority_arn",
		max:           200,
		min:           5,
		pattern:       regexp.MustCompile(`^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:[\w+=/,.@-]*:[0-9]*:[\w+=,.@-]+(/[\w+=,.@-]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule) Name() string {
	return "aws_acmpca_certificate_authority_certificate_invalid_certificate_authority_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAcmpcaCertificateAuthorityCertificateInvalidCertificateAuthorityArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"certificate_authority_arn must be 200 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"certificate_authority_arn must be 5 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:[\w+=/,.@-]*:[0-9]*:[\w+=,.@-]+(/[\w+=,.@-]+)*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}