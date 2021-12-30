package user

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"index"`
	Nickname string
	Password string
	Token    string
	Room     int
}

func (this Player) SetNickame(name string) {
	this.Nickname = name
}
