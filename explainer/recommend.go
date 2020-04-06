package explainer

import "fmt"

type recmd struct {
	messages []string
}

func NewRecmd() Recommend{
	return &recmd{}
}

type Recommend interface {
	GetAllRecommend() []string
	Recommend(t int, name string)
	HasRecommend() bool
}

func (r *recmd) Recommend(t int, name string){
	m := fmt.Sprintf("%s: %s is not valid.", getTypeText(t), name)
	r.messages = append(r.messages, m)
}

func (r *recmd) GetAllRecommend() []string{
	return r.messages
}

func (r *recmd) HasRecommend() bool{
	return r.messages != nil
}
