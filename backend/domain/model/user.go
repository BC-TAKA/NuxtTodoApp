package model

type User struct {
	ID   int    `xorm:"id"`
	Name string `xorm:"name"`
	Todo string `xorm:"todo"`
}
