package model

type Test struct {
	Id        int   `gorm:"column:id;primary_key"`
	Value     int   `gorm:"column:value"`
	CreatedAt int64 `gorm:"column:created_at"`
}

/**
测试数据表
*/
func (e *Test) TableName() string {
	return "test"
}
