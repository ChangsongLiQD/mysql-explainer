package explainer

import "strings"

const (
	RULE_SELECT_TYPE = iota
	RULE_TYPE
	RULE_EXTRA
)

type rule struct {
	st map[string]bool
	t  map[string]bool
	extra map[string]bool
}

func NewRule() *rule{
	st := make(map[string]bool)
	t := make(map[string]bool)
	extra := make(map[string]bool)
	return &rule{st, t, extra}
}

type Rule interface {
	AddRule(t int, n string, valid bool)
	CheckRule(t int, n string) bool
}

func (r *rule) AddRule(t int, n string, valid bool){
	switch t {
	case RULE_SELECT_TYPE:
		r.st[n] = valid
	case RULE_TYPE:
		r.t[n] = valid
	case RULE_EXTRA:
		r.extra[n] = valid
	}
}

func (r *rule) CheckRule(t int, n string) bool{
	var valid, ok bool
	switch t {
	case RULE_SELECT_TYPE:
		valid, ok = r.st[n]
	case RULE_TYPE:
		valid, ok = r.t[n]
	case RULE_EXTRA:
		return r.checkExtra(n)
	}

	if !ok {
		return true
	}
	return valid
}

func (r *rule) checkExtra(n string) bool{
	if n == "" {
		return true
	}

	for rule := range r.extra {
		find := strings.Contains(n, rule)
		if find{
			return false
		}
	}
	return true
}

func getTypeText(t int) string{
	switch t {
	case RULE_SELECT_TYPE:
		return "Select Type"
	case RULE_TYPE:
		return "Type"
	case RULE_EXTRA:
		return "Extra"
	}
	return "未知类型"
}