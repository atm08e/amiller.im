package tests

import (
	"github.com/revel/revel/testing"
	"amiller.im/app/models"
	"log"
 	"github.com/davecgh/go-spew/spew"
	//"time"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}


func (t *AppTest) TestPostModel(){
	log.Println("TestPostModel")
	models.InitDatabaseConnection()
	models.InsertNewPost("TestAuthor1","TestTitle1", "subTitle", "TestBody1", []models.Tag{})
/*
	post := models.GetPostById(1)
	//log.Println("Retrieving Test Post: ")
	//log.Println(post)
	spew.Dump(post)

	log.Println("Trying to updating post!")
	models.UpdatePostBody(1, "UPDATED POST BODY")
	postUpdated := models.GetPostById(1)
	log.Println("Retrieving Updated Test Post: ")
	log.Println(postUpdated)
	spew.Dump(postUpdated)
	*/
	// Delete Post
	//log.Println("Trying to delete post post_id=1")
	//models.DeletePost(1)

	// Find All Posts
	log.Println(("Trying to Select * (all)"))
	posts := models.GetAllPost()
	spew.Dump(posts)

}