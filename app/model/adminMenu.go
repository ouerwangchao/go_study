package model

import "time"

type AdminMenu struct {
	Id        int    `json:"id"`
	ParentId  int    `json:"parent_id"`
	Sort      int    `json:"sort"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Uri       string `json:"uri"`
	Display   int    `json:"display"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

/**
后台菜单表
*/
func (e *AdminMenu) TableName() string {
	return "admin_menu"
}
