package model

type BlogType struct {
	BlogType string `gorm:"column:blog_type" json:"blog_type"`
}
type BlogTypeCategory struct {
	BlogType string `gorm:"column:blog_type" json:"blog_type"`
	Count    string `gorm:"column:count(*)" json:"count(*)"`
}
