package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name        string
	Technique   string
	Base        string
	Ingredients string
}

type Ingredient struct {
	Name   string
	Amount float64
}
