---
title: Charmil Validator
slug: /charmil_validator
---

Validator can be used for testing and controlling many aspects of cobra commands. It provides many rules out of the box for validating the commands.

## Rules provided by validator
- LengthRule
- MustExistRule
- UseMatchesRule
> We are working on the validator library to provide more rules

## How to use
It is recommended to use the validator while writing unit tests for cobra commands.

1. Create a configuration of type `rules.ValidatorConfig`. You can provide your own ValidatorConfig or use the default one by leaving it empty
```go
var ruleCfg rules.ValidatorConfig
```
or overriding default config
```go
ruleCfg := rules.ValidatorConfig{
	ValidatorRules: rules.ValidatorRules{
		Length: rules.Length{Limits: map[string]rules.Limit{"Use": {Min: 1}}},
		MustExist: rules.MustExist{Fields: map[string]bool{"Args": true}},
		UseMatches: rules.UseMatches{Regexp: `^[^-_+]+$`},
	},
}
```
2. Generate the validation errors by using `ExecuteRules` function over the ruleCfg
```go
validationErr := rules.ExecuteRules(cmd, &ruleCfg)
```
`ExecuteRules` function will return a slice of `ValidationError` object, which can be efficiently used for testing purposes.
```go
if len(validationErr) != 0 {
	t.Errorf("validationErr was not empty, got length: %d; want %d", len(validationErr), 0)
}
for _, errs := range validationErr {
	if errs.Err != nil {
		t.Errorf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
	}
}
```

## Ignore Commands
Sometimes during development, you want to pass the tests for certain commands, but at the same time use Validator for tests. Validation can be skipped/ignored for the commands, mentioned in the validator configuration.
To ignore the commands you need to specify the path of the command in validator configuration.

 Skip single command `mycli actions create`
```go
ValidatorOptions: rules.ValidatorOptions{
	SkipCommands: map[string]bool{"mycli actions create": true},
},
```

2. Skip the command including all children
```go
ValidatorOptions: rules.ValidatorOptions{
	SkipChildren: map[string]bool{"mycli": true},
},
```

3. Skip the command for specific rule
```go
Length: rules.Length{
	RuleOptions: validator.RuleOptions{
		SkipCommands: map[string]bool{"mycli actions create": true},
	},
	Limits: map[string]rules.Limit{
		"Use": {Min: 1},
	},
},
