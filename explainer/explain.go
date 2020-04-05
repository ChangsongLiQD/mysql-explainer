package explainer

type Explain struct {
	Id int `gorm:"id" json:"id"`
	SelectType string `gorm:"select_type" json:"select_type"`
	Table string `gorm:"table" json:"table"`
	Type string `gorm:"type" json:"type"`
	PossibleKey string `gorm:"possible_keys" json:"possible_keys"`
	Key string `gorm:"key" json:"key"`
	KeyLen int `gorm:"key_len" json:"key_len"`
	Ref string `gorm:"ref" json:"ref"`
	Rows int `gorm:"rows" json:"rows"`
	Extra string `gorm:"Extra" json:"Extra"`
}
