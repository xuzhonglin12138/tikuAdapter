// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameTiku = "tiku"

// Tiku mapped from table <tiku>
type Tiku struct {
	ID           int32  `gorm:"type:integer primary key autoincrement" json:"id"`
	Question     string `gorm:"column:question;type:longtext" json:"question"`
	QuestionText string `gorm:"column:question_text;type:longtext;comment:只保留汉子数字字母，方便模糊匹配" json:"question_text"` // 只保留汉子数字字母，方便模糊匹配
	Type         int32  `gorm:"column:type;type:int(11)" json:"type"`
	Options      string `gorm:"column:options;type:longtext" json:"options"`
	Answer       string `gorm:"column:answer;type:longtext" json:"answer"`
	Plat         int32  `gorm:"column:plat;type:int(11)" json:"plat"`
	QuestionHash string `gorm:"column:question_hash;type:char(16);index:question_hash_idx,priority:1;comment:只有问题的短hash" json:"question_hash"` // 只有问题的短hash
	Hash         string `gorm:"column:hash;type:char(32);uniqueIndex:hash_index,priority:1;comment:整个实体的hash,防止重复" json:"hash"`                // 整个实体的hash,防止重复
	Source       int32  `gorm:"column:source;type:int(11);comment:0采集1自建2文件类" json:"source"`                                                   // 0采集1自建2文件类
	Extra        string `gorm:"column:extra;type:text;comment:扩展字段,多用于tag" json:"extra"`                                                       // 扩展字段,多用于tag
}

// TableName Tiku's table name
func (*Tiku) TableName() string {
	return TableNameTiku
}
