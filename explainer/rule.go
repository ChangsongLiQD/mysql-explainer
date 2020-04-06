package explainer

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
	return &rule{}
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
		valid, ok = r.extra[n]
	}

	if !ok {
		return true
	}
	return valid
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