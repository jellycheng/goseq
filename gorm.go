package goseq

type GormCommonField struct {
	ID         int64 `gorm:"primary_key;Column:id" json:"id"`
	IsDelete   int   `gorm:"Column:is_delete" json:"is_delete"`
	CreateTime int64 `gorm:"Column:create_time" json:"create_time"`
	UpdateTime int64 `gorm:"Column:update_time" json:"update_time"`
	DeleteTime int64 `gorm:"Column:delete_time" json:"delete_time"`
}
