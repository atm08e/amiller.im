package models

type Tag struct {
	TagId   uint 	`gorm:"primary_key"`
	Name 	string 	`sql:"size:64;unique"`
	//Posts 	[]Post 	`gorm:"many2many:posts_to_tags;"`
}