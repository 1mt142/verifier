package models

type Article struct {
	ID         uint
	Title      string
	Content    string
	CategoryID uint
	Category   Category
	Tags       []*Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	ID       uint
	Name     string
	Articles []*Article `gorm:"many2many:article_tags;"`
}

type Category struct {
	ID       uint
	Name     string
	Articles []Article
}
