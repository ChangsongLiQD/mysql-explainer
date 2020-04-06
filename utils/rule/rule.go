package rule

import "mysql-explainer/explainer"

func GetRule(ruleConfig map[string] []string) explainer.Rule{
	rule := explainer.NewRule()
	for t, ruleNames := range ruleConfig {
		tp := getTypeNum(t)
		for _, name := range ruleNames {
			rule.AddRule(tp, name, false)
		}
	}
	return rule
}

func getTypeNum(t string) int{
	switch t {
	case "selecttype":
		return explainer.RULE_SELECT_TYPE
	case "type":
		return explainer.RULE_TYPE
	default:
		return explainer.RULE_EXTRA
	}
}