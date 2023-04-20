package common

type GeneralList[T any] struct {
	Page int64 `json:"page"`
	Data T     `json:"data"`
}

type PageList[T any] struct {
	Current  int    ` json:"current"`
	PageSize int    `json:"pageSize"`
	Expr     Search `json:"expr"`
	Page     int64  `json:"page"`
	Data     T      `json:"data"`
	startAt  int
	endAt    int
	sel      string
}

type PageChange struct {
	Current  int    ` json:"current"`
	PageSize int    `json:"pageSize"`
	Expr     Search `json:"expr"`
}

type Search struct {
	Picker   []string `json:"picker"`
	Valve    bool     `json:"valve"`
	Text     string   `json:"text"`
	Explain  string   `json:"explain"`
	WorkId   string   `json:"work_id"`
	Type     int      `json:"type"`
	Status   int      `json:"status"`
	IDC      string   `json:"idc"`
	Source   string   `json:"source"`
	Username string   `json:"username"`
	Dept     string   `json:"dept"`
	RealName string   `json:"real_name"`
	Rule     string   `json:"rule"`
	Email    string   `json:"email"`
	IP       string   `json:"ip"`
	IsQuery  int      `json:"is_query"`
}

type SQLTest struct {
	SourceId string `json:"source_id"`
	SQL      string `json:"sql"`
	Database string `json:"data_base"`
	Kind     int    `json:"kind"`
	WorkId   string `json:"work_id"`
}

type QueryOrder struct {
	SourceId string `json:"source_id"`
	Export   uint   `json:"export"`
	Text     string `json:"text"`
	WorkId   string `json:"work_id"`
	Tp       string `json:"tp"`
}

type _dbInfo struct {
	Results   []string
	QueryList []map[string]interface{}
}

type FieldInfo struct {
	Field      string  `gorm:"Column:Field" json:"field"`
	Type       string  `gorm:"Column:Type" json:"type"`
	Collation  string  `gorm:"Column:Collation" json:"collation"`
	Null       string  `gorm:"Column:Null" json:"null"`
	Key        string  `gorm:"Column:Key" json:"key"`
	Default    *string `gorm:"Column:Default" json:"default"`
	Extra      string  `gorm:"Column:Extra" json:"extra"`
	Privileges string  `gorm:"Column:Privileges" json:"privileges"`
	Comment    string  `gorm:"Column:Comment" json:"comment"`
}

type IndexInfo struct {
	Table      string `gorm:"Column:Table"`
	NonUnique  int    `gorm:"Column:Non_unique"`
	IndexName  string `gorm:"Column:Key_name"`
	Seq        int    `gorm:"Column:Seq_in_index"`
	ColumnName string `gorm:"Column:Column_name"`
	IndexType  string `gorm:"Column:Index_type"`
}

type Resp struct {
	Payload interface{} `json:"payload"`
	Code    int         `json:"code"`
	Text    string      `json:"text"`
}

const (
	ORDER_IS_CREATE     = "工单已创建!"
	ORDER_IS_DUP        = "工单请勿重复提交!"
	ORDER_IS_EDIT       = "工单已编辑！"
	ORDER_IS_DELETE     = "工单已删除！"
	ORDER_IS_CLEAR      = "工单已清除"
	ORDER_IS_AGREE      = "工单已同意"
	ORDER_IS_UNDO       = "工单已撤销"
	ORDER_IS_REJECT     = "工单已拒绝"
	ORDER_IS_ALL_END    = "所有工单已终止"
	ORDER_IS_END        = "工单已终止"
	ORDER_IS_ALL_CANCEL = "所有工单已取消"
	DATA_IS_DELETE      = "数据已删除！"
	DATA_IS_EDIT        = "数据已编辑！"
	DATA_IS_UPDATED     = "数据已更新"
)

const (
	Pong = "pong"
	Ping = "ping"
	PING = 1
)
