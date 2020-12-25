package model

type User struct {
	ID   int    `xorm:"id"`
	Name string `xorm:"name"`
	Todo string `xorm:"todo"`
}

type InsertParams struct {
	Name string `json:"name" xorm:"name"`
	Todo string `json:"todo" xorm:"todo"`
}
