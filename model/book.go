package model

type Book struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(100)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
	PublishDate string `gorm:"type:date" json:"publishDate"`
}
