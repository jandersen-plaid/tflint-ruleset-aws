// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDxBgpPeerInvalidAddressFamilyRule checks the pattern is valid
type AwsDxBgpPeerInvalidAddressFamilyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDxBgpPeerInvalidAddressFamilyRule returns new rule with default attributes
func NewAwsDxBgpPeerInvalidAddressFamilyRule() *AwsDxBgpPeerInvalidAddressFamilyRule {
	return &AwsDxBgpPeerInvalidAddressFamilyRule{
		resourceType:  "aws_dx_bgp_peer",
		attributeName: "address_family",
		enum: []string{
			"ipv4",
			"ipv6",
		},
	}
}

// Name returns the rule name
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Name() string {
	return "aws_dx_bgp_peer_invalid_address_family"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as address_family`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}