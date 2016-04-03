package models
import(
	"time"
	"log"
)


var db = GetDatabaseConnection()

//type Post interface {
//	GetTemplatedBody() template.HTML
//}

// TODO change Created/Updated these back to datetime objects
type Post struct {
	PostId 		uint 		`gorm:"primary_key"`
	Author 		string 		`sql:"size:64"`
	Title 		string 		`sql:"size:64"`
	Subtitle 	string 		`sql:"size:256"`
	Body 		string 		`sql:"type:longtext"`
	Tags 		[]Tag 		`gorm:"many2many:posts_to_tags;"`
	Created 	string		`sql:"size:64"`
	Updated 	string		`sql:"size:64"`
}

func InsertNewPost(author string, title string, subtitle string, body string, tags []Tag){
	log.Println("InsertNewPost")
	//db := GetDatabaseConnection()
	newPost := Post{
		Author: author,
		Title: title,
		Subtitle: subtitle,
		Body: body,
		Tags: tags,
		Created: time.Now().Format(time.RFC1123),
		Updated: time.Now().Format(time.RFC1123),
	}
	result := db.NewRecord(newPost)
	if !result{
		log.Println("The records primary key is not blank (Error)")
	}else{
		// TODO More Meta
		log.Println("The records primary key is blank (OK)")
	}
	db.Create(&newPost)
}

func UpdatePostBody(id uint, body string){
	post:= GetPostById(id)
	post.Body = body
	post.Updated = time.Now().Format(time.RFC1123)
	//db := GetDatabaseConnection()
	db.Save(&post)
}

func GetLastNPost(n int) []Post{
	posts := []Post{}
	db.Limit(n).Find(&posts)
	return posts
}

func GetAllPost() []Post{
	//dbconn  := GetDatabaseConnection()
	posts := []Post{}
	db.Find(&posts)
	return posts
}

func DeletePost(id uint){
	//db := GetDatabaseConnection()
	deletePost := Post{
		PostId: id,
	}
	db.Delete(&deletePost)
}

func GetPostById(id uint) Post{
	//db := GetDatabaseConnection()
	post := Post{}
	db.Where("post_id = ?", id).First(&post)
	return post
}