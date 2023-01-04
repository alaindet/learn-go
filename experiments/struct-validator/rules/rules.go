package rules

type Rule struct {
	name string
	err  error
}

type RuleInterface interface {
	Name() string
	Error() string
	Run(val any)
}

func (r *Rule) Name() string {
	return r.name
}

func (r *Rule) Error() string {
	if r.err == nil {
		return ""
	}
	return r.err.Error()
}
