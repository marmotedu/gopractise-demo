package main

import (
	"time"

	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	AnimalID int64     `gorm:"column:animalID"` // 将列名设为 `animalID`
	Birthday time.Time `gorm:"column:birthday"` // 将列名设为 `birthday`
	Age      int64     `gorm:"column:age"`      // 将列名设为 `age`
}

func (a *Animal) TableName() string {
	return "animal"
}
