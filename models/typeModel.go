package models

import "gorm.io/gorm"

type TypeTree struct {
	gorm.Model
	Name     string      `json:"name"`
	ParentID *uint       `json:"-"`
	Parent   *TypeTree   `json:"parent" gorm:"foreignkey:ParentID"`
	Children []*TypeTree `json:"children" gorm:"foreignkey:ParentID"`
}
