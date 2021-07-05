package rules

import (
	"encoding/json"

	"github.com/aerogear/charmil/validator"
	"github.com/imdario/mergo"
	"github.com/spf13/cobra"
)

// ValidatorConfig is provided to user for overriding default rules
type ValidatorConfig struct {
	Options        `json:"Options"`
	ValidatorRules `json:"ValidatorRules"`
}
type Options struct {
	Verbose bool `json:"Verbose"`
}
type ValidatorRules struct {
	Length    `json:"Length"`
	MustExist `json:"MustExist"`
}

// ExecuteRules executes all the rules
// provided by validatorConfig
func (validatorConfig *ValidatorConfig) ExecuteRules(cmd *cobra.Command) []validator.ValidationError {
	var ruleConfig RuleConfig
	return ruleConfig.ExecuteRulesInternal(cmd, validatorConfig)
}

func ValidatorConfigToRuleConfig(validatorConfig *ValidatorConfig, ruleConfig *RuleConfig) {

	defaultConfigJson := `{
		"Options": {
			"Verbose": false
		},
		"ValidatorRules": {
			"Length": {
				"Limits": {
					"Use":     {"Min": 2},
					"Short":   {"Min": 15},
					"Long":    {"Min": 50},
					"Example": {"Min": 50}
				}
			},
			"MustExist": {
				"Fields": {"Use": true, "Short": true, "Long": true, "Example": true}
			}
		}
	}`

	// unmarshal defaultConfigJson in configHelper
	var configHelper ValidatorConfig
	json.Unmarshal([]byte(defaultConfigJson), &configHelper)

	// Merge user provided config into configHelper
	mergo.Merge(&configHelper, validatorConfig, mergo.WithSliceDeepCopy)
	validatorConfig = &configHelper

	// append rules to execute
	ruleConfig.Verbose = validatorConfig.Verbose
	ruleConfig.Rules = append(ruleConfig.Rules, &validatorConfig.Length)
	ruleConfig.Rules = append(ruleConfig.Rules, &validatorConfig.MustExist)
}
