package rules

type Rule struct {
	Name string
	Err  error
}

type RuleInterface interface {
	Run(val any)
}
