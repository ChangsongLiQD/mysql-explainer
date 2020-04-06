package explainer

type optimize struct {
	exp Explain
	r   Rule
}

func NewOptimize() Optimizer{
	return &optimize{}
}

type Optimizer interface {
	SetExplain(exp Explain)
	SetRule(r Rule)
	CheckRule() Recommend
}

func (o *optimize) SetExplain(exp Explain){
	o.exp = exp
}

func (o *optimize) SetRule(r Rule){
	o.r = r
}

func (o *optimize) CheckRule() Recommend{
	m := map[int]string{
		RULE_SELECT_TYPE: o.exp.SelectType,
		RULE_TYPE: o.exp.Type,
		RULE_EXTRA: o.exp.Extra,
	}

	rec := NewRecmd()
	for t, n := range m{
		valid := o.r.CheckRule(t, n)
		if !valid {
			rec.Recommend(t, n)
		}
	}
	return rec
}
