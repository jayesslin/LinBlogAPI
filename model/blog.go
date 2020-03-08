package model

type AdditionBlog struct {
	BlogTitle   string `gorm:"column:blog_title" json:"blog_title"`
	BlogContent string `gorm:"column:blog_content" json:"blog_content"`
}
type Blog struct {
	BlogID      int    `gorm:"column:id" json:"id"`
	BlogTitle   string `gorm:"column:blog_title" json:"blog_title,omitempty"`
	BlogContent string `gorm:"column:blog_content" json:"blog_content,omitempty"`
	BlogCreate  string `gorm:"column:create_time" json:"create_time,omitempty"`
	BlogUpdate  string `gorm:"column:update_time" json:"update_time,omitempty"`
	BlogAuthor  string `gorm:"column:blog_author" json:"blog_author,omitempty"`
}
