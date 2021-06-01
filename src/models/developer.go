package models

import "gorm.io/gorm"

type Developer struct {
	gorm.Model `swaggerignore:"true"`
	Name       string
	Website    string
}
