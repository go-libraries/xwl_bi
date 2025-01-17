package model

type App struct {
	Id         int    `db:"id" json:"id"`
	AppName    string `db:"app_name" json:"app_name"`
	Descibe    string `db:"descibe" json:"descibe"`
	AppId      string `db:"app_id" json:"app_id"`
	AppKey     string `db:"app_key" json:"app_key"`
	CreateBy   int    `db:"create_by" json:"create_by"`
	UpdateBy   int    `db:"update_by" json:"update_by"`
	CreateTime string `db:"create_time" json:"create_time"`
	UpdateTime string `db:"update_time" json:"update_time"`
	AppManager string `db:"app_manager" json:"app_manager"`
	SaveMonth  int    `db:"save_mouth" json:"save_mouth"`
	IsClose    *int   `db:"is_close" json:"is_close"`
	Page       uint64 `json:"page" db:"-"`
	Limit      uint64 `json:"limit" db:"-"`
}
