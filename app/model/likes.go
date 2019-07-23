package model

import "time"

type Likes struct {
	Id        int       `gorm:"column:id;primary_key"`
	Ip        string    `gorm:"column:ip"`
	Ua        string    `gorm:"column:ua"`
	Title     string    `gorm:"column:title"`
	Hash      string    `gorm:"column:hash"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

// TableName sets the insert table name for this struct type
func (e *Likes) TableName() string {
	return "likes"
}
