package explainer

// select type
const (
	ST_SIMPLE   = "SIMPLE"
	ST_PARIMARY = "PRIMARY"
	ST_SUBQUERY = "SUBQUERY"
	ST_DERIVED = "DERIVED"
	ST_UNION = "UNION"
	ST_UNION_RESULT = "UNION RESULT"
)

// type
const (
	T_ALL = "ALL"
	T_INDEX = "index"
	T_RANGE = "range"
	T_REF = "ref"
	T_EQ_REF = "eq_ref"
	T_CONST = "const"
	T_SYSTEM = "system"
)

// extra
const (
	X_FILESORT = "Using filesort"
	X_TEMPORARY = "Using temporary"
	X_INDEX = "Using index"
	X_WHERE = "Using where"
	X_JOIN_BUFFER = "Using join buffer"
	X_IMPOSSIBLE_WHERE = "impossible where"
	X_OPTIMIZED_AWAY = "select tables optimized away"
	X_DISTINCT = "distinct"
)

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